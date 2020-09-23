package routers

import (
	"daidai-server/middleware/jwt"
	"daidai-server/pkg/setting"
	"daidai-server/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	gApi := r.Group("/api")
	gApi.Use(jwt.JWT())

	apiDemand := gApi.Group("/demand")

	{
		apiDemand.POST("/", api.AddDemand)
		apiDemand.GET("/homeRecommend", api.HomeRecommend)
		apiDemand.GET("/page", api.Page)
	}

	apiOss := gApi.Group("/oss")
	{
		apiOss.POST("/upload", api.UploadFile)
	}

	apiUser := gApi.Group("/user")
	{
		apiUser.POST("/regist", api.AddUsers)
		apiUser.POST("/login", api.Login)
		apiUser.GET("/getToken", api.GetToken)
		apiUser.POST("/setUserInfo", api.SetUserInfo)
	}
	return r
}
