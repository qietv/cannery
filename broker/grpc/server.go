package grpc

import (
	"github.com/qietv/qgrpc"
	"google.golang.org/grpc"
)

var (
	grpcServer qgrpc.Server
)

type server struct {
}

func Startup() {
	s, err := qgrpc.New(&qgrpc.Config{
		Name:        "qietv",
		Network:     "tcp",
		Addr:        ":8808",
		AccessLog:   "access.log",
		ErrorLog:    "error.log",
		Interceptor: nil,
	}, func(s *grpc.Server) {

	})
	if err != nil {
		panic("grpc server start fail")
	}
	defer s.Server.GracefulStop()
}
