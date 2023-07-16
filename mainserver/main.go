package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"growth/conf"
	"growth/dbhelper"
	"growth/handle"
	"growth/initialize"
	"growth/pb"
)

func initConf() {
	// default UTC time location
	time.Local = time.UTC
	// Load global config
	conf.LoadConfigs()
	// Initialize Logger
	initialize.InitLogger()
	// Initialize db
	dbhelper.InitDb()
}

func main() {
	initConf()

	lis, err := net.Listen("tcp", ":80")
	if err != nil {
		zap.S().Errorf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterUserCoinServer(server, &handle.UgCoinServer{})
	pb.RegisterUserGradeServer(server, &handle.UgGradeServer{})

	reflection.Register(server)

	zap.S().Infof("server listening at %v\n", lis.Addr())

	//启动服务
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
