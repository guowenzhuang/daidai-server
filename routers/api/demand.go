package api

import (
	"daidai-server/models"
	"daidai-server/pkg/setting"
	"daidai-server/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

//新增需求
func AddDemand(c *gin.Context) {
	var demand models.Demand //声明一个UserInfo类型的变量u
	err := c.ShouldBind(&demand)
	if err != nil {
		log.Println("新增需求错误:", err)
	}
	// 获取当前用户
	claimsInterface, _ := c.Get(setting.CurrentUserInfo)
	claims := claimsInterface.(*util.Claims)
	demand.CreateUserId = claims.Id

	models.AddDemand(demand)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "新增成功",
		"data": nil,
	})
}

func HomeRecommend(c *gin.Context) {
	recommend := models.HomeRecommend()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": recommend,
	})
}

func Page(c *gin.Context) {
	pageStr, limitStr, name, order := c.Query("page"), c.Query("limit"), c.Query("name"), c.Query("order")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	data := models.Page((page-1)*limit, limit, name, order)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "查询成功",
		"data": data,
	})
}
