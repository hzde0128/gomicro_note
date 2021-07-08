package handlers

import (
	"context"
	"gomicro_note/p20/models"

	"github.com/gin-gonic/gin"
)

// PanicIfError panic异常处理
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

// GetProdDetail 显示商品详情
func GetProdDetail(c *gin.Context) {
	var prodReq models.ProdRequest
	PanicIfError(c.BindUri(&prodReq))
	prodService := c.Keys["prodservice"].(models.ProdService)
	res, _ := prodService.GetProdDetail(context.Background(), &prodReq)

	c.JSON(200, gin.H{"data": res.Data})
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
		prodRes, _ := prodService.GetProdList(context.Background(), &prodReq)

		c.JSON(200, gin.H{"data": prodRes.Data})

		// 超时代码
		// 1.配置config
		//configA := hystrix.CommandConfig{
		//	Timeout: 1000,
		//}
		//// 2.配置command
		//hystrix.ConfigureCommand("getProds", configA)
		//// 3.执行Do方法
		//var prodRes *models.ProdListResponse
		//err := hystrix.Do("getProds", func() error {
		//	prodRes, err = prodService.GetProdList(context.Background(), &prodReq)
		//	return err
		//}, func(e error) error {
		//	// 降级 显示默认产品
		//	prodRes, err = defaultProds()
		//	return err
		//})
	}

}
