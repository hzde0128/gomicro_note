package main

import "strconv"

// ProdModel 商品模型
type ProdModel struct {
	// ProdId 商品Id
	ProdID int `json:"prod_id"`
	// ProdName 商品名称
	ProdName string `json:"prod_name"`
}

// NewProd 新增产品
func NewProd(id int, name string) *ProdModel {
	return &ProdModel{ProdID: id, ProdName: name}
}

// NewProdList 产品列表
func NewProdList(n int) []*ProdModel {
	ret := make([]*ProdModel, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "Prod"+strconv.Itoa(100+i)))
	}
	return ret
}
