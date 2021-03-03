package grpc

import (
	"cannery/broker/api"
	"context"
	"github.com/qietv/qgrpc"
)

var (
	grpcServer qgrpc.Server
)

type server struct {}


func Startup() {
	s, err := qgrpc.New(&qgrpc.Config{
		Name:        "cannery",
		Network:     "tcp",
		Addr:        ":8808",
		AccessLog:   "/var/log/cannery/access.log",
		ErrorLog:    "/var/log/cannery/error.log",
		Interceptor: nil,
	}, func(gs *qgrpc.Server) {
		api.RegisterBrokerServer(gs.Server, &server{})
	})
	if err != nil {
		panic("grpc server start fail")
	}
	defer s.Server.GracefulStop()
}




func (s *server) Register(context.Context, *api.RegisterReq) (*api.RegisterResp, error) {
	panic("implement me")
}

func (s *server) Heartbeat(context.Context, *api.HeartbeatReq) (*api.HeartbeatResp, error) {
	panic("implement me")
}

func (s *server) Unregister(context.Context, *api.UnRegisterReq) (*api.UnRegisterResp, error) {
	panic("implement me")
}

func (s *server) Commit(context.Context, *api.CommitReq) (*api.CommitResp, error) {
	panic("implement me")
}