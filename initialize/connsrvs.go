package initialize

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"growth/global"
	"growth/pb"
)

func Conn_srv() {
	// 连接到grpc服务的客户端
	conn, err := grpc.Dial("localhost:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	global.ClientCoin = pb.NewUserCoinClient(conn)
	global.ClientGrade = pb.NewUserGradeClient(conn)
}
