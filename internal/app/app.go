package app

import (
	grpc_app "github.com/pedroxer/resource-service/internal/app/grpc"
	"github.com/pedroxer/resource-service/internal/services/items"
	"github.com/pedroxer/resource-service/internal/services/workplace"
	"github.com/pedroxer/resource-service/internal/storage"
	log "github.com/sirupsen/logrus"
)

type App struct {
	GRPCSrv *grpc_app.App
}

func NewApp(log *log.Logger, grpcPort int, store *storage.Storage) *App {
	wrkpl := workplace.NewDefaultWorkplaceService(store, log)
	items := items.NewDefaultItemService(store, log)
	grpcApp := grpc_app.NewApp(log, grpcPort, wrkpl, items)

	return &App{
		GRPCSrv: grpcApp,
	}
}
