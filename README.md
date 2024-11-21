# AiChatPartner-api

## 环境
1. go 1.23.2 linux/amd64
2. go-zero v1.7.3
3. goctl 1.7.3 linux/amd64

## 运行

```shell
# api server
cd api
go mod tidy
go run api.go -f etc/api-api.yaml

# rpc server
cd rpc
go mod tidy
go run rpc.go -f etc/rpc.yaml
```
