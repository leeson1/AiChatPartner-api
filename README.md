<!--
 * @Author: Leeson
 * @Date: 2024-11-24 00:37:56
-->
# AiChatPartner-api

## 环境
1. go 1.23.2 linux/amd64
2. goctl 1.7.3 linux/amd64
3. mysql 5.7.43

## 运行

```shell
# api server
cd api
go mod tidy
go run api.go -f etc/api-api.yaml

# rpc server
cd rpc/chat
go mod tidy
go run chat.go -f etc/chat.yaml
```
