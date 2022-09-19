# GRPC 开发模板

# 依赖
* go 1.18.5
* grpc 1.49.0
* zap 1.23.0 (日志)
# 工具
* tproxy 报文调试 [https://github.com/kevwan/tproxy]
* protobuf 3.21.5 [https://github.com/protocolbuffers/protobuf/releases]


# 安装插件
1. go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
2. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# proto 生成代码
protoc -I . *.proto --go_out=. --go-grpc_out=. 

# 环境变量
`GODEBUG=http2debug=2` 调试

# 待解决问题

* [x] 心跳保活
  * MaxConnectionAge 设置为最大值