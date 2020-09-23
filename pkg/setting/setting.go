package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg             *ini.File
	RunMode         string
	HTTPPort        int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	PageSize        int
	JwtSecret       string
	AppId           string
	AppSecret       string
	CurrentUserInfo string
)

func init() {
	CurrentUserInfo = "currentUserInfo"

	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("读取 conf/app.ini 错误 %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
	LoadOss()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("读取项目配置失败 :%v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

func LoadOss() {
	sec, err := Cfg.GetSection("wx")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	AppId = sec.Key("appId").MustString("")
	AppSecret = sec.Key("appSecret").MustString("")

}
