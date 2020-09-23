package models

import (
	"daidai-server/pkg/util"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Users struct {
	gorm.Model
	// 用户名称
	Name string
	// 用户密码
	Password string
	// 用户手机号
	Phone string
	// 用户邮箱
	Email string
	// 是否锁定
	Locking bool
	// openId
	OpenId string
	// 昵称
	NickName string
	// 头像
	AvatarUrl string
	// 性别 1:男 2:女
	Gender uint
	// 国家
	Country string
	// 省份
	Province string
	// 城市
	City string
}

func AddUser(users *Users) bool {
	encodePassword := util.EncryptPassword(users.Password)
	users.Password = encodePassword
	DB.Create(&users)
	return true
}

func Login(account string, password string) (bool, Users) {
	users := Users{}
	DB.Where("name = ?", account).First(&users)
	if (users == Users{}) {
		DB.Where("phone = ?", account).First(&users)
	}
	if (users == Users{}) {
		DB.Where("email = ?", account).First(&users)
	}

	if (users == Users{}) {
		return false, Users{}
	}
	return util.MatchPassword(password, users.Password), users
}

func GetToken(code, appId, appSecret string) (users Users, openId, sessionKey string) {
	Url, _ := url.Parse("https://api.weixin.qq.com/sns/jscode2session")
	params := url.Values{}
	params.Set("appid", appId)
	params.Set("secret", appSecret)
	params.Set("js_code", code)
	params.Set("grant_type", "authorization_code")
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	log.Println("请求获取openId", urlPath) //等同于https://www.xxx.com?age=23&name=zhaofan
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	wxRespMap := make(map[string]string)
	json.Unmarshal(body, &wxRespMap)
	users = Users{}
	DB.Where("open_id = ?", wxRespMap["openid"]).First(&users)
	return users, wxRespMap["openid"], wxRespMap["session_key"]
}
