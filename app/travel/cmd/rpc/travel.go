package main

import (
	"flag"
	"fmt"
	"home-nest/pkg/interceptor/rpcserver"

	"home-nest/app/travel/cmd/rpc/internal/config"
	"home-nest/app/travel/cmd/rpc/internal/server"
	"home-nest/app/travel/cmd/rpc/internal/svc"
	"home-nest/app/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/travel.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterTravelServer(grpcServer, server.NewTravelServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
