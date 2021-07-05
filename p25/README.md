# 使用micro api进行接口测试

运行micro api

```bash
#!/bin/bash

export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
export MICRO_API_HANDLER=rpc
export MICRO_API_NAMESPACE=api.hzde.com

micro api
```

发送POST请求

```bash
curl -X POST "http://localhost:8080/test/TestService/Call" -d '{"id": 1}'
```

输出结果

```json
{"data":"test1"}
```
