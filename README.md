# 项目工具

## 自动生成 gRPC 文件

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative growth.proto
# protoc --go_out=:. --go-grpc_out=require_unimplemented_servers=false:. -I . growth.proto
```

## xorm 生成数据模型

````bash
go install xorm.io/reverse@latest

cd database
reverse -f mysql-growth.yml
````

## grpcurl 调用

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

## 生成 protoset 文件

```bash
protoc --proto_path=. --descriptor_set_out=myservice.protoset --include_imports ./growth.proto
```

## 生成 grpc-gateway 代码

````bash
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

protoc -I . --grpc-gateway_out ./ \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    growth.proto
````

## 客户端的http请求

````bash
# gin路由
curl -X 'GET' "localhost:8080/v1/Growth.UserCoin/ListTasks"
curl -X 'GET' "localhost:8080/v1/Growth.UserGrade/ListGrades"
# grpc-gateway代理
curl -X 'GET' "localhost:8081/v1/Growth.UserCoin/ListTasks"
curl -X 'GET' "localhost:8081/v1/Growth.UserGrade/ListGrades"

````

## 加入Prometheus统计

```bash
# ginserver
# 代码文件 /comm/prometheus.go
# 头部定义数据模型
# var metricsRequest prometheus.Gauge

# 在 main.go 中加入相应方法
# main方法中创建和注册数据以及监听 /metrics路径
#     MetricInit

# 服务方法调用的中间件中加入统计指标
#     MetricAdd

# 测试请求
curl -X'GET' "localhost:8080/v1/Growth.UserCoin/ListTasks"
# 访问数据
curl "localhost:8080/metrics"
```

# 构建部署
在项目根目录下

## 构建 growth-system 服务程序

```bash
CGO_ENABLED=0&& GOOS=linux&& GOARCH=AMD64&& go build -o growth_client mainclient/main.go
CGO_ENABLED=0&& GOOS=linux&& GOARCH=AMD64&& go build -o growth_server mainserver/main.go
CGO_ENABLED=0&& GOOS=linux&& GOARCH=AMD64&& go build -o growth_api ginserver/main.go

mv growth_* dockerfile/
```

## 构建 grpcurl 工具

```bash
git clone git@github.com:fullstorydev/grpcurl.git
cd ./grpcurl/cmd/grpcurl
CGO_ENABLED=0&& GOOS=linux&& GOARCH=AMD64&& go build .

mv grpcurl.exe dockerfile/
```

## 构建 docker 文件

```bash
docker build . --file=./dockerfile/Dockerfile --network=host --platform=linux/amd64 -t growthsystem:v1.0
```

## 本地启动容器

```bash
docker run --network=host -it growthsystem:v1.0 /bin/bash

# 查看容器
docker container ls
docker exec -it container_id /bin/bash
```

## 推送到 TCR 服务端

```bash
docker login -u username -p password
docker push growthsystem:v1.0
```

# K8S 配置部署

## 拉取镜像(work节点)

```bash
docker login -u username -p password
docker pull growthsystem:v1.0
```

## 创建deployment(master节点)

```bash
# 查看帮助 kubectl create deployment --help
kubectl create namespace k8stest
kubectl create deployment -n k8stest growthsystem --image=/k8stest/growthsystem:v1.0 --replicas=1
```
## 创建service(master节点)

```bash
# 查看帮助
kubectl create service --help
# 创建指定namespace的服务
# 将容器内的80端口暴露为80端口对外提供服务
kubectl create service clusterip -n k8stest growthsystem --tcp=80:80
kubectl create service clusterip -n k8stest growthsystem --tcp=8080:8080
kubectl create service clusterip -n k8stest growthsystem --tcp=8081:8081
# 删除一个service
kubectl delete service -n k8stest growthsystem
# 将一个deployment暴露为service:
kubectl expose deployment -n k8stest growthsystem --port=80 --target-port=80 --cluster-ip=
kubectl get svc -A -o wide
```

## 查看部署、pod的信息(master节点)

```bash
kubectl get deployment -n k8stest
kubectl get pods -n k8stest -o wide
kubectl logs -n k8stest growthsystem-xxxx-xxxx
kubectl describe pod growthsystem-xxxx-xxxx
```

## 验证 growthsystem 程序

```bash
docker container ls
docker exec -it container_id /bin/bash
# pod内部本地访问
./growthclient --addr=127.0.0.1:80
# pod的ip地址
./growthclient --addr=10.244.1.3:80 
# 同命名空间的pod内访问service名称
./growthclient --addr=growthsystem:80
# service域名
./growthclient --addr=growthsystem.k8stest.svc.cluster.local:80
```