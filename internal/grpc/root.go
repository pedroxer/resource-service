package grpc

import (
	"github.com/pedroxer/resource-service/internal/proto_gen"
	"google.golang.org/grpc"
)

type serverAPI struct {
	proto_gen.UnimplementedResourceServiceServer
	items      ItemService
	workplaces WorkplaceService
	parkings   ParkingService
}

func Register(gRPC *grpc.Server, itemsService ItemService, service WorkplaceService, parkings ParkingService) {
	proto_gen.RegisterResourceServiceServer(gRPC, &serverAPI{items: itemsService, workplaces: service, parkings: parkings})
}
