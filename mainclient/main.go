package main

import (
	"context"
	"flag"
	"sync"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"growth/initialize"
	"growth/pb"
)

var connPool = sync.Pool{
	New: func() any {
		// 连接到服务
		addr := flag.String("addr", "localhost:80", "the address to connect to")
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithWriteBufferSize(1024 * 1024 * 1), // 默认32KB
			grpc.WithReadBufferSize(1024 * 1024 * 1),  // 默认32KB,
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                10 * time.Minute,
				Timeout:             10 * time.Second,
				PermitWithoutStream: false,
			}),
		}
		conn, err := grpc.Dial(*addr, opts...)
		if err != nil {
			zap.S().Errorf("did not connect: %v", err)
		}
		return conn
	},
}

func GetConn() *grpc.ClientConn {
	return connPool.Get().(*grpc.ClientConn)
}
func CloseConn(conn *grpc.ClientConn) {
	connPool.Put(conn)
}

func main() {
	initialize.InitLogger()

	conn := GetConn()
	if conn != nil {
		defer CloseConn(conn)
	} else {
		zap.S().Errorf("connection nil")
	}
	// 请求服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 新建客户端
	cCoin := pb.NewUserCoinClient(conn)
	// 测试1：UserCoinServer.ListTasks
	r1, err1 := cCoin.ListTasks(ctx, &pb.ListTasksRequest{})
	if err1 != nil {
		zap.S().Infof("cCoin.ListTasks error=%v\n", err1)
	} else {
		zap.S().Infof("cCoin.ListTasks: %+v\n", r1.GetDatalist())
	}
}
