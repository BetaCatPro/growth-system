# 自动生成 gRPC 文件

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative growth.proto
# protoc --go_out=:. --go-grpc_out=require_unimplemented_servers=false:. -I . growth.proto
```

# xorm 生成数据模型

````bash
go install xorm.io/reverse@latest

cd database
reverse -f mysql-growth.yml
````

# grpcurl 调用

```bash
# 安装 grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# 使用gRPC服务
grpcurl -plaintext localhost:80 list
grpcurl -plaintext localhost:80 list Growth.UserCoin
grpcurl -plaintext localhost:80 describe
grpcurl -plaintext localhost:80 describe Growth.UserCoin
grpcurl -plaintext localhost:80 describe Growth.UserCoin.ListTasks
# 使用proto文件
grpcurl -import-path ./ -proto growth.proto list
# 使用 protoset 文件
grpcurl -protoset myservice.protoset list Growth.UserCoin
# 调用gRPC服务
grpcurl -plaintext localhost:80 Growth.UserCoin/ListTasks
grpcurl -plaintext -d '{"uid":1}' localhost:80 Growth.UserCoin/UserCoinInfo
```

# 生成 protoset 文件

```bash
protoc --proto_path=. --descriptor_set_out=myservice.protoset --include_imports ./growth.proto
```

# 生成 grpc-gateway 代码

````bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

protoc -I . --grpc-gateway_out ./ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    growth.proto
````

# 客户端的http请求

````bash
# gin路由
curl -X 'GET' "localhost:8080/v1/Growth.UserCoin/ListTasks"
curl -X 'GET' "localhost:8080/v1/Growth.UserGrade/ListGrades"
# grpc-gateway代理
curl -X 'GET' "localhost:8081/v1/Growth.UserCoin/ListTasks"
curl -X 'GET' "localhost:8081/v1/Growth.UserGrade/ListGrades"

````

# 加入Prometheus统计

````
代码文件 prometheus.go
头部定义数据模型
var metricsRequest prometheus.Gauge

在main.go中加入相应方法
main方法中创建和注册数据以及监听/metrics路径
    MetricInit

服务方法调用的中间件中加入统计指标
    MetricAdd

测试请求
    curl -X'GET' "localhost:8080/v1/Growth.UserCoin/ListTasks"
访问数据
    curl "localhost:8080/metrics"
````
