package middlewares

import (
	"gomicro_note/p17/grpc_client/models"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(prodService models.ProdService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Keys = make(map[string]interface{})
		c.Keys["prodservice"] = prodService
	}
}
