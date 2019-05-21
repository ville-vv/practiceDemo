# grpc-gateway 的使用

gRPC 是一个高性能、通用的开源RPC框架，其由 Google 主要面向移动应用开发并基于HTTP/2 协议标准而设计，基于 ProtoBuf(Protocol Buffers) 序列化协议开发，且支持众多开发语言。

使用 grpc+protobuf 通信在许多微服务上都可以看到他们的身影。不仅可以使各个模块独立。还可以使用不同的编程语言服务互通。但是单纯的 grpc+protobuf 也存在一定的限制，两个服务间，或者客户端与服务间必须都要使用 protobuf 数据解析。

grpc-gateway + protobuf 增加了服务的功能，服务之间不仅可以使用 rpc 通信，还可以可以使用 web API 通信。当前这个 practice 是使用 grpc-gateway 的示例。

要启动两个程序，一个用来做http服务代理，一个用来做grpc 通信。

- proxy 里面的代码是http服务代理
- simple 是 grpc+protobuf 协议代码
- main.go grpc 服务

http 代理服务地址 0.0.0.0:32112
grpc 服务地址 0.0.0.0:32111

http 访问示例

```
curl --request POST \
  --url http://localhost:32112/v1/echo \
  --header 'cache-control: no-cache' \
  --header 'content-type: application/json' \
  --header 'postman-token: abd73bb6-f9bf-c888-3e76-5892cf63df6d' \
  --data '{\n	"cmd":3333,\n	"message":"server grpc http proxy",\n	"timeStamp":222222222224444444,\n	"token":"JEKGJCHESDFRGDDDDDDDDDDDDDDDDDDD",\n	"msgType":1,\n	"body":"5o6l5Y+j57un5om/IHByb3RvYnVm55qEZ3JwYyDkuK3nmoQgU2ltcGxlU2VydmVyU2VydmVyIOaOpeWPo+OAgiDor7fmsYLmlbDmja7lhaXlj6PmmK/lnKjov5nph4w="\n}'
```
