package jwt

import (
	"daidai-server/pkg/setting"
	"daidai-server/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type FilterUrl struct {
	Url, Method string
}

var permission = [10]FilterUrl{{
	Url:    "/api/oss/upload",
	Method: "POST",
}, {
	Url:    "/api/user/regist",
	Method: "POST",
}, {
	Url:    "/api/user/login",
	Method: "POST",
}, {
	Url:    "/api/user/getToken",
	Method: "GET",
}, {
	Url:    "/api/user/setUserInfo",
	Method: "POST",
}, {
	Url:    "/api/demand/homeRecommend",
	Method: "GET",
},
	{
		Url:    "/api/demand/page",
		Method: "GET",
	},
	{
		Url:    "/api/user/setUserInfo",
		Method: "POST",
	},
	{
		Url:    "/api/user/regist",
		Method: "POST",
	}}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.Request
		path := request.URL.Path
		method := request.Method
		for i := 0; i < len(permission); i++ {
			item := permission[i]
			if path == item.Url && method == item.Method {
				c.Next()
				return
			}
		}

		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			returnFail(c, "token未传")
			return
		}

		token := authorization[7:len(authorization)]
		if token == "" {
			returnFail(c, "token未获取到")
			return
		}
		log.Printf("请求url:%s , token:%s", path, token)
		claims, err := util.ParseToken(token)
		if err != nil {
			returnFail(c, "token校验失败")
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			returnFail(c, "token已过期")
			return
		}
		c.Set(setting.CurrentUserInfo, claims)
		c.Next()
	}

}

func returnFail(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 502,
		"msg":  msg,
		"data": nil,
	})
	c.Abort()
}
