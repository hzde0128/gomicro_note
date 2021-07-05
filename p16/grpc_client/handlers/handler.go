package handlers

import (
	"context"
	"gomicro_note/p16/grpc_client/models"

	"github.com/gin-gonic/gin"
)

// GetProdList 显示商品列表
func GetProdList(c *gin.Context) {
	prodService := c.Keys["prodservice"].(models.ProdService)
	var prodReq models.ProdRequest
	err := c.Bind(&prodReq)
	if err != nil {
		c.JSON(500, gin.H{
			"status": err.Error()})
		return
	}
	if prodReq.Size == 0 {
		prodReq.Size = 2
	}
	prodRes, _ := prodService.GetProdList(context.Background(), &prodReq)
	c.JSON(200, gin.H{
		"data": prodRes.Data,
	})

}
