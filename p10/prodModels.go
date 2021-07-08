package main

import "strconv"

// ProdModel 商品模型
type ProdModel struct {
	ProdID   int    `json:"prod_id"`
	ProdName string `json:"prod_name"`
}

// ProdsRequest 请求信息打tag
type ProdsRequest struct {
	Size int `form:"size"`
}

// NewProd 新增商品
func NewProd(id int, name string) *ProdModel {
	return &ProdModel{ProdID: id, ProdName: name}
}

// NewProdList 根据前端请求的size返回商品数量
func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "Prod"+strconv.Itoa(100+i)))
	}
	return ret
}
