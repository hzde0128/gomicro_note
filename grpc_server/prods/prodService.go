package prods

import (
	"context"
	"grpc_server/models"
	"strconv"
)

// 服务实现
type ProdService struct {}

func newProd(id int32, pname string) *models.ProdModel{
	return &models.ProdModel{ProdId:id, ProdName:pname}
}

//GetProdList(ctx context.Context, in *ProdRequest, out *ProdListResponse) error
func (*ProdService) GetProdList(ctx context.Context,in *models.ProdRequest,res *models.ProdListResponse) error{
	models := make([]*models.ProdModel,0)
	var i int32
	for i=0;i<in.Size;i++{
		models = append(models, newProd(100 +i, "prodName"+ strconv.Itoa(100+int(i))))
	}
	res.Data = models
	return nil
}