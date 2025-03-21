package server

import (
	v1 "github.com/BitofferHub/seckill/api/sec_kill/proto"
	"github.com/BitofferHub/seckill/internal/conf"
	"github.com/BitofferHub/seckill/internal/service"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer
//
//	@Author <a href="https://bitoffer.cn">狂飙训练营</a>
//	@Description: NewGRPCServer new a gRPC server.
//	@param c
//	@param greeter
//	@return *grpc.Server
func NewGRPCServer(c *conf.Server, greeter *service.SecKillService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			mmd.Server(),
			MiddlewareTraceID(),
			MiddlewareLog(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterSecKillServer(srv, greeter)
	return srv
}
