package grpc

import (
	"github.com/pedroxer/resource-service/internal/proto_gen"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type serverAPI struct {
	proto_gen.UnimplementedResourceServiceServer
	logger     *log.Logger
	items      ItemService
	workplaces WorkplaceService
	parkings   ParkingService
}

func Register(gRPC *grpc.Server, log *log.Logger, itemsService ItemService, service WorkplaceService, parkings ParkingService) {
	proto_gen.RegisterResourceServiceServer(gRPC, &serverAPI{items: itemsService, workplaces: service, parkings: parkings, logger: log})
}
