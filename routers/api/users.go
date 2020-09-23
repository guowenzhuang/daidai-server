package api

import (
	"daidai-server/models"
	"daidai-server/pkg/setting"
	"daidai-server/pkg/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func SetUserInfo(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	setUserReqpMap := make(map[string]string)
	json.Unmarshal(data, &setUserReqpMap)
	resp, err := util.Dnctypt(setUserReqpMap["encryptedData"], setUserReqpMap["sessionKey"], setUserReqpMap["iv"])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "解谜失败",
			"data": nil,
		})
		log.Println("解谜失败:", err)
		return
	}

	log.Println("解谜成功", resp)

	userInfoMap := make(map[string]string)
	json.Unmarshal([]byte(resp), &userInfoMap)
	gender, _ := strconv.ParseUint(userInfoMap["gender"], 10, 0)
	// 设置用户信息
	user := models.Users{
		Name:      userInfoMap["nickName"],
		Password:  "qq123456",
		Phone:     "",
		Email:     "",
		Locking:   false,
		OpenId:    userInfoMap["openId"],
		NickName:  userInfoMap["nickName"],
		AvatarUrl: userInfoMap["avatarUrl"],
		Gender:    uint(gender),
		Country:   userInfoMap["country"],
		Province:  userInfoMap["province"],
		City:      userInfoMap["city"],
	}
	models.AddUser(&user)
	token, err := util.GenerateToken(user.ID, user.NickName, user.AvatarUrl)
	if err != nil {
		log.Println("获取token失败：", err)
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "获取token失败",
			"data": nil,
		})
		return
	}
	respData := make(map[string]string)
	respData["token"] = token
	respData["nickName"] = user.NickName
	respData["avatarUrl"] = user.AvatarUrl

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "设置用户信息成功",
		"data": respData,
	})
}

func GetToken(c *gin.Context) {
	code := c.Query("code")
	user, _, sessionKey := models.GetToken(code, setting.AppId, setting.AppSecret)

	if (user == models.Users{}) {
		// 没有用户信息
		respData := make(map[string]string)
		respData["sessionKey"] = sessionKey

		c.JSON(http.StatusOK, gin.H{
			"code": 300,
			"msg":  "没有用户信息",
			"data": respData,
		})
	} else {
		// 获取token
		token, err := util.GenerateToken(user.ID, user.NickName, user.AvatarUrl)
		if err != nil {
			log.Println("获取token失败：", err)
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "获取token失败",
				"data": nil,
			})
			return
		}
		respData := make(map[string]string)
		respData["token"] = token
		respData["nickName"] = user.NickName
		respData["avatarUrl"] = user.AvatarUrl

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "获取用户信息成功",
			"data": respData,
		})
	}
}

func AddUsers(c *gin.Context) {
	var users models.Users //声明一个UserInfo类型的变量u
	err := c.ShouldBind(&users)
	if err != nil {
		log.Println("新增用户错误:", err)
	}
	models.AddUser(&users)
	fmt.Print(users.ID)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "新增成功",
		"data": nil,
	})
}

func Login(c *gin.Context) {
	account, password := c.Query("account"), c.Query("password")
	isVerificationSuccess, user := models.Login(account, password)
	if isVerificationSuccess {
		// 获取token
		token, err := util.GenerateToken(user.ID, account, password)
		if err != nil {
			log.Println("获取token失败：", err)
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "获取token失败",
				"data": nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "登陆成功",
			"data": token,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "密码错误",
			"data": nil,
		})
	}

}
