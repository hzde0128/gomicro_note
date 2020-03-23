package handlers

import (
	"context"
	"gomicro_note/p18/grpc_client/models"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func newProd(id int32, pname string) *models.ProdModel {
	return &models.ProdModel{ProdId: id, ProdName: pname}
}

func defaultProds() (*models.ProdListResponse, error) {
	model := make([]*models.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		model = append(model, newProd(20+i, "prodName"+strconv.Itoa(20+int(i))))
	}
	res := &models.ProdListResponse{}
	res.Data = model
	return res, nil
}

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
			Timeout: 1000,
		}
		// 2.配置command
		hystrix.ConfigureCommand("getProds", configA)
		// 3.执行Do方法
		var prodRes *models.ProdListResponse
		err := hystrix.Do("getProds", func() error {
			prodRes, err = prodService.GetProdList(context.Background(), &prodReq)
			return err
		}, func(e error) error {
			// 降级 显示默认产品
			prodRes, err = defaultProds()
			return err
		})
		if err != nil {
			c.JSON(500, gin.H{"status": err.Error()})
		} else {
			c.JSON(200, gin.H{"data": prodRes.Data})
		}

	}

}
