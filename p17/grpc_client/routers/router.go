package routers

import (
	"gomicro_note/p17/grpc_client/handlers"
	"gomicro_note/p17/grpc_client/middlewares"
	"gomicro_note/p17/grpc_client/models"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由
func InitRouter(prodService models.ProdService) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.InitMiddleware(prodService))
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", handlers.GetProdList)
	}
	return r
}
