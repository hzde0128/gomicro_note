package routers

import (
	"gomicro_note/p20/grpc_client/handlers"
	"gomicro_note/p20/grpc_client/middlewares"
	"gomicro_note/p20/grpc_client/models"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由
func InitRouter(prodService models.ProdService) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.InitMiddleware(prodService), middlewares.ErrorMiddleware())
	v1Group := r.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", handlers.GetProdList)
		v1Group.Handle("GET", "/prods/:pid", handlers.GetProdDetail)
	}
	return r
}
