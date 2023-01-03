package helper

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kekasicoid/kekasigohelper"
	"github.com/spf13/viper"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
)

var (
	limit   ratelimit.Limiter
	limiter = rate.NewLimiter(1, 3)
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	limit = ratelimit.New(1) // 1 request per second
	return &GoMiddleware{}
}

func (m *GoMiddleware) LimitRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.Allow() == false {
			c.AbortWithStatusJSON(429, http.StatusTooManyRequests)
			return
		}
		return
	}
}

func (m *GoMiddleware) LeakBucketUber() gin.HandlerFunc {
	prev := time.Now()
	return func(c *gin.Context) {
		now := limit.Take()
		kekasigohelper.LoggerInfo(now.Sub(prev))
		prev = now
	}
}

func (m *GoMiddleware) GinMiddlewareSocketIo(allowOrigin string) gin.HandlerFunc {
	envMode := viper.Get("APP_MODE").(string)
	return func(c *gin.Context) {
		if envMode != "development" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Request.Header.Del("Origin")
		c.Header("Content-Type", gin.MIMEJSON)
		c.Next()
	}
}
