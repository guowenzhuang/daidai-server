package api

import (
	"daidai-server/models"
	"daidai-server/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddUsers(c *gin.Context) {
	var users models.Users //声明一个UserInfo类型的变量u
	err := c.ShouldBind(&users)
	if err != nil {
		log.Println("新增用户错误:", err)
	}
	models.AddUser(users)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "新增成功",
		"data": nil,
	})
}

func Login(c *gin.Context) {
	account, password := c.Query("account"), c.Query("password")
	isVerificationSuccess := models.Login(account, password)
	if isVerificationSuccess {
		// 获取token
		token, err := util.GenerateToken(account, password)
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
