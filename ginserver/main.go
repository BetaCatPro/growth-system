package main

import (
	"context"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	"growth/comm"
	"growth/conf"
	"growth/dbhelper"
	"growth/global"
	"growth/handle"
	"growth/initialize"
	"growth/pb"
	"growth/router"
)

func initConf() {
	// default UTC time location
	time.Local = time.UTC
	// Load global config
	conf.LoadConfigs()
	// Initialize Logger
	initialize.InitLogger()
	// Initialize grpc server connections
	initialize.Conn_srv()
	// Initialize db
	dbhelper.InitDb()
}

func mainGateway() {
	s := grpc.NewServer()
	// 注册服务
	pb.RegisterUserCoinServer(s, &handle.UgCoinServer{})
	pb.RegisterUserGradeServer(s, &handle.UgGradeServer{})
	reflection.Register(s)

	// grpc-gateway 注册服务
	serveMuxOpt := []runtime.ServeMuxOption{
		runtime.WithOutgoingHeaderMatcher(func(s string) (string, bool) {
			return s, true
		}),
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {
			origin := request.Header.Get("Origin")
			if global.AllowOrigin[origin] {
				md := metadata.New(map[string]string{
					"Access-Control-Allow-Origin":      origin,
					"Access-Control-Allow-Methods":     "GET,POST,PUT,DELETE,OPTION",
					"Access-Control-Allow-Headers":     "*",
					"Access-Control-Allow-Credentials": "true",
				})
				grpc.SetHeader(ctx, md)
			}
			return nil
		}),
	}
	mux := runtime.NewServeMux(serveMuxOpt...)
	ctx := context.Background()
	if err := pb.RegisterUserCoinHandlerServer(ctx, mux, &handle.UgCoinServer{}); err != nil {
		zap.S().Infof("Faile to RegisterUserCoinHandlerServer error=%v", err)
	}
	if err := pb.RegisterUserGradeHandlerServer(ctx, mux, &handle.UgGradeServer{}); err != nil {
		zap.S().Infof("Faile to RegisterUserGradeHandlerServer error=%v", err)
	}
	httpMux := http.NewServeMux()
	httpMux.Handle("/v1/Growth", mux)
	// 配置http服务
	server := &http.Server{
		Addr: ":8081",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			zap.S().Infof("http.HandlerFunc url=%s", r.URL)
			mux.ServeHTTP(w, r)
		}),
	}
	// 启动http服务
	zap.S().Infof("server.ListenAdnServe(%s)", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		zap.S().Fatalf("ListenAndServe error=%v", err)
	}
}

func mainGin() {
	// prometheus client Create non-global registry.
	router := router.GetRouters()
	comm.MetricInit(router)

	// 为http/2配置参数
	h2Handler := h2c.NewHandler(router, &http2.Server{})
	// 配置http服务
	server := &http.Server{
		Addr:    ":8080",
		Handler: h2Handler,
	}
	// 启动http服务
	server.ListenAndServe()
}

func main() {
	initConf()
	go mainGateway()
	mainGin()
}
