package app

import (
	grpc_app "github.com/pedroxer/resource-service/internal/app/grpc"
	"github.com/pedroxer/resource-service/internal/storage"
	log "github.com/sirupsen/logrus"
)

type App struct{
	GRPCSrv *grpc_app.App
}

func NewApp(log *log.Logger, grpcPort int, store *storage.Storage) *App{
	grpcApp := grpc_app.NewApp(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
	
}
