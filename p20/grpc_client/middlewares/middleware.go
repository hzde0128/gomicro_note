package middlewares

import (
	"fmt"
	"gomicro_note/p20/models"

	"github.com/gin-gonic/gin"
)

// InitMiddleware 注入prodService中间件
func InitMiddleware(prodService models.ProdService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["prodservice"] = prodService
	}
}

// ErrorMiddleware 异常处理
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(500, gin.H{"status": fmt.Sprintf("%s", r)})
				c.Abort()
			}
		}()
		c.Next()
	}
}
