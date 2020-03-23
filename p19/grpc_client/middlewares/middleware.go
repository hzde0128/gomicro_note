package middlewares

import (
	"gomicro_note/p19/grpc_client/models"

	"github.com/gin-gonic/gin"
)

// InitMiddleware 注入prodService中间件
func InitMiddleware(prodService models.ProdService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["prodservice"] = prodService
	}
}
