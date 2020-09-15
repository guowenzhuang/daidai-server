package routers

import (
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

	apiDemand := gApi.Group("/demand")

	{
		apiDemand.POST("/", api.AddDemand)
	}
	return r
}
