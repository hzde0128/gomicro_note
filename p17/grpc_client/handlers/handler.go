package handlers

import (
	"context"
	"gomicro_note/p17/models"

	"github.com/afex/hystrix-go/hystrix"
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
	} else {
		// 超时代码
		// 1.配置config
		configA := hystrix.CommandConfig{
			Timeout: 5000,
		}
		// 2.配置command
		hystrix.ConfigureCommand("getProds", configA)
		// 3.执行Do方法
		var prodRes *models.ProdListResponse
		err := hystrix.Do("getProds", func() error {
			prodRes, err = prodService.GetProdList(context.Background(), &prodReq)
			return err
		}, nil)
		if err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			c.JSON(200, gin.H{"data": prodRes.Data})
		}

	}

}
