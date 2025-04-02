package grpc

import (
	"context"
	"github.com/pedroxer/resource-service/internal/models"
	"github.com/pedroxer/resource-service/internal/proto_gen"
	"github.com/pedroxer/resource-service/internal/utills"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type ParkingService interface {
	GetParkingSpaces(ctx context.Context, address, zone, spaceType string, isAvailable bool, page int64) ([]models.ParkingPlace, int64, error)
	GetParkingSpaceById(ctx context.Context, id int64) (models.ParkingPlace, error)
	CreateParkingSpace(ctx context.Context, parkingSpace models.ParkingPlace) (int64, error)
	UpdateParkingSpace(ctx context.Context, parkingSpace models.ParkingPlace) (models.ParkingPlace, error)
	DeleteParkingSpace(ctx context.Context, id int64) error
}

func (s *serverAPI) GetParkingSpaces(ctx context.Context, req *proto_gen.GetParkingSpacesRequest) (*proto_gen.GetParkingSpacesResponse, error) {
	if req.GetPage() == 0 {
		req.Page = 1
	}
	parkingSpaces, amount, err := s.parkings.GetParkingSpaces(ctx, req.Address, req.Zone, req.Type, req.IsAvailable, req.Page)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := new(proto_gen.GetParkingSpacesResponse)
	response.ParkingSpaces = make([]*proto_gen.ParkingSpace, 0)
	for _, parkingSpace := range parkingSpaces {
		response.ParkingSpaces = append(response.ParkingSpaces, castServiceParkingSpaceToProto(parkingSpace))
	}
	response.PageSize = utills.PageSize
	response.TotalCount = amount/utills.PageSize + 1
	response.Page = req.Page
	return response, nil
}

func (s *serverAPI) GetParkingSpaceById(ctx context.Context, req *proto_gen.GetParkingSpaceByIdRequest) (*proto_gen.ParkingSpace, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	parkingSpace, err := s.parkings.GetParkingSpaceById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return castServiceParkingSpaceToProto(parkingSpace), nil
}

func (s *serverAPI) CreateParkingSpace(ctx context.Context, req *proto_gen.CreateParkingSpaceRequest) (*proto_gen.ParkingSpace, error) {
	if req.Address == "" || req.Zone == "" || req.Type == "" || req.Number == 0 {
		return nil, status.Error(codes.InvalidArgument, "address, zone, number and type are required")
	}

	id, err := s.parkings.CreateParkingSpace(ctx, models.ParkingPlace{
		Number:      req.Number,
		Address:     req.Address,
		Type:        req.Type,
		IsAvailable: req.IsAvailable,
		Zone:        req.Zone,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return castServiceParkingSpaceToProto(models.ParkingPlace{
		Id:          id,
		Number:      req.Number,
		Address:     req.Address,
		Type:        req.Type,
		IsAvailable: req.IsAvailable,
		Zone:        req.Zone,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}), nil
}

func (s *serverAPI) UpdateParkingSpace(ctx context.Context, req *proto_gen.UpdateParkingSpaceRequest) (*proto_gen.ParkingSpace, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	parkingSpace, err := s.parkings.UpdateParkingSpace(ctx, models.ParkingPlace{
		Id:          req.Id,
		Number:      req.Number,
		Address:     req.Address,
		Type:        req.Type,
		IsAvailable: req.IsAvailable,
		Zone:        req.Zone,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return castServiceParkingSpaceToProto(parkingSpace), nil
}
func (s *serverAPI) DeleteParkingSpace(ctx context.Context, req *proto_gen.DeleteParkingSpaceRequest) (*proto_gen.DeleteParkingSpaceResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	err := s.parkings.DeleteParkingSpace(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.DeleteParkingSpaceResponse{Success: true}, nil
}

func castServiceParkingSpaceToProto(parkingSpace models.ParkingPlace) *proto_gen.ParkingSpace {
	return &proto_gen.ParkingSpace{
		Id:          parkingSpace.Id,
		Number:      parkingSpace.Number,
		Address:     parkingSpace.Address,
		Type:        parkingSpace.Type,
		IsAvailable: parkingSpace.IsAvailable,
		Zone:        parkingSpace.Zone,
		CreatedAt:   timestamppb.New(parkingSpace.CreatedAt),
		UpdatedAt:   timestamppb.New(parkingSpace.UpdatedAt),
	}
}
