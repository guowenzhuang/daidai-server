package api

import (
	"daidai-server/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//新增需求
func AddDemand(c *gin.Context) {
	var demand models.Demand //声明一个UserInfo类型的变量u
	err := c.ShouldBind(&demand)
	if err != nil {
		log.Println("新增需求错误:", err)
	}
	models.AddDemand(demand)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "新增成功",
		"data": nil,
	})
}
