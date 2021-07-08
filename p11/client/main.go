package main

import (
	"context"
	"gomicro_note/p11/models"
	"log"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	http "github.com/micro/go-plugins/client/http/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

// consul 通过轮询获取服务
// 调用http api 引入protobuf生成请求响应模型
func callAPI(s selector.Selector) {
	myClient := http.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("ProdSrv", "/v1/prods",
		models.ProdRequest{Size: 6})
	var rsp models.ProdListResponse
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(rsp.GetData())
}

func main() {
	// consul 连接句柄
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"))

	sel := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	callAPI(sel)

}
