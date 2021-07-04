package main

import (
	"context"
	"fmt"

	http "github.com/asim/go-micro/plugins/client/http/v3"
	etcd "github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/selector"
)

// 由于http的插件还是1.0的，其它的client还是使用v1版本
// etcd 通过轮询获取服务
// 使用插件 调用http api
func callAPI(s selector.Selector) (map[string]interface{}, error) {
	myClient := http.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	// 请求对象
	req := myClient.NewRequest("ProdSrv", "/v1/prods", map[string]string{})
	// 响应对象
	var rsp map[string]interface{}
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func main() {
	// etcd 连接句柄
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))

	sel := selector.NewSelector(
		selector.Registry(etcdReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	resp, err := callAPI(sel)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Println(resp["data"])
}
