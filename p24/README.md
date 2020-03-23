# 第24讲:使用Micro工具查看和调用我们的服务

## 环境变量设置

```bash
export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
```

## 获取服务的详细信息

```bash
micro get service test.hzde.com
```

输出结果

```bash
service  test.hzde.com

version latest

ID	Address	Metadata
test.hzde.com-7e8c75dd-75e9-49a3-a415-48dbfad893ac	[fdf8:292:d5e9:0:ccc:f210:43a8:3f2b]:8000	transport=grpc,broker=eats,protocol=grpc,registry=etcd,server=grpc

Endpoint: TestService.Call

Request: {
	id int32
}

Response: {
	data string
}
```

## 通过micro工具箱请求后端服务

```bash
micro call test.hzde.com TestService.Call '{"id": 123}'
```

输出结果

```bash
{
	"data": "test123"
}
```

