package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
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
