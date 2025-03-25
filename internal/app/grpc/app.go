package grpc_app

import (
	"fmt"
	"net"

	mygrpc "github.com/pedroxer/resource-service/internal/grpc"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type App struct {
	logger     *log.Logger
	grpcServer *grpc.Server
	port       int
}

func NewApp(log *log.Logger, port int, workplaceSerivce mygrpc.WorkplaceService) *App {
	server := grpc.NewServer()
	mygrpc.Register(server, nil, workplaceSerivce, nil)
	return &App{
		logger:     log,
		grpcServer: server,
		port:       port,
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatal(err)
	}
	a.logger.Infof("starting grpc server on port %d", a.port)
	return a.grpcServer.Serve(l)
}

func (a *App) Stop() {
	a.logger.Info("stopping grpc server")
	a.grpcServer.GracefulStop()
}
