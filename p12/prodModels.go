package p12

import "strconv"

type ProdModel struct {
	ProdId int `json:"pid"`
	ProdName string `json:"pname"`
}

type ProdsRequest struct {
	Size int `form:"size"`
}

func NewProd(id int, name string) *ProdModel{
	return &ProdModel{ProdId:id, ProdName:name}
}

func NewProdList(n int) []*ProdModel{
	ret := make([]*ProdModel,0)
	for i:=0 ; i<n ; i++ {
		ret = append(ret, NewProd(100+i, "Prod"+strconv.Itoa(i)))
	}
	return  ret
}