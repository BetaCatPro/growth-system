# 自动生成 gRPC 文件

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative growth.proto
# protoc --go_out=:. --go-grpc_out=require_unimplemented_servers=false:. -I . growth.proto
```

# 安装 xorm，生成数据模型

````bash
cd database
reverse -f mysql-growth.yml
````
