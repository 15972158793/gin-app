package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// RateLimit 令牌桶限流 1s多少个
func RateLimit(fillInterval time.Duration, capacity int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, capacity)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": -3,
				"msg":  "已被限流",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
