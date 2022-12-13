package middleware

import (
	"net/http"
	"strings"

	"github.com/15972158793/gin-app/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "user_id"
const ContextUserLogin = "login"

func JWTAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -2,
				"msg":  "Authorization不存在",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authorization, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": -2,
				"msg":  "Authorization格式错误",
			})
			c.Abort()
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -2,
				"msg":  "Token失效",
			})
			c.Abort()
			return
		}

		// 根据生成的token和redis存储的对比可以实现同时只可以登录一个账号

		c.Set(ContextUserIDKey, mc.UserID)
		c.Next()
	}
}
