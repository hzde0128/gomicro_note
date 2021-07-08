package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	http "github.com/micro/go-plugins/client/http/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

// 由于http的插件还是1.0的，其它的client还是使用v1版本
// consul 通过轮询获取服务
// 使用插件 调用http api
func callAPI(s selector.Selector) {
	myClient := http.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("ProdSrv", "/v1/prods", map[string]string{})
	var rsp map[string]interface{}
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(rsp["data"])
}

func main() {
	// consul连接句柄
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"))

	sel := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	callAPI(sel)

}
