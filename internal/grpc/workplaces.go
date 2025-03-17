package grpc

import (
	"context"
	"github.com/pedroxer/resource-service/internal/models/services"
	"github.com/pedroxer/resource-service/internal/proto_gen"
	"github.com/pedroxer/resource-service/internal/utills"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type WorkplaceService interface {
	GetWorkplaces(zone, floor, workplaceType string, capacity int64, isAvailable bool, page, pageSize int64) ([]services.Workplace, int64, error)
	GetWorkplacesById(id int64) (services.Workplace, error)
	CreateWorkplace(workplace services.Workplace) (services.Workplace, error)
	UpdateWorkplace(workplace services.Workplace) (services.Workplace, error)
	DeleteWorkplace(id int64) error
	CheckWorkplaceAvailability(id int64) (bool, []services.TimeSlot, error)
}

func (s *serverAPI) GetWorkplaces(ctx context.Context, req *proto_gen.GetWorkplacesRequest) (*proto_gen.GetWorkplacesResponse, error) {
	if req.GetPage() == 0 {
		req.Page = 1
	}
	workplaces, amount, err := s.workplaces.GetWorkplaces(req.Zone, req.Floor, req.Type, req.Capacity, req.IsAvailable, req.Page, utills.PageSize)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := new(proto_gen.GetWorkplacesResponse)
	response.Workplaces = make([]*proto_gen.Workplace, 0)
	for _, workplace := range workplaces {
		response.Workplaces = append(response.Workplaces, castServiceWorkplaceToProto(workplace))
	}
	response.PageSize = utills.PageSize
	response.TotalCount = amount / utills.PageSize
	response.Page = req.Page
	return response, nil
}

func (s *serverAPI) GetWorkplaceById(ctx context.Context, req *proto_gen.GetWorkplaceByIdRequest) (*proto_gen.Workplace, error) {
	workplace, err := s.workplaces.GetWorkplacesById(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return castServiceWorkplaceToProto(workplace), nil
}
func (s *serverAPI) CreateWorkplace(ctx context.Context, req *proto_gen.CreateWorkplaceRequest) (*proto_gen.Workplace, error) {
	workplace, err := s.workplaces.CreateWorkplace(services.Workplace{
		Address:           req.Address,
		Zone:              req.Zone,
		Floor:             req.Floor,
		Number:            req.Number,
		Type:              req.Type,
		Capacity:          req.Capacity,
		Description:       req.Description,
		IsAvailable:       req.IsAvailable,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		MaintenanceStatus: req.MaintenanceStatus,
		UniqueTag:         req.UniqueTag,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return castServiceWorkplaceToProto(workplace), nil
}

func (s *serverAPI) UpdateWorkplace(ctx context.Context, req *proto_gen.UpdateWorkplaceRequest) (*proto_gen.Workplace, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	workplace, err := s.workplaces.UpdateWorkplace(services.Workplace{
		Id:                req.Id,
		Address:           req.Address,
		Zone:              req.Zone,
		Floor:             req.Floor,
		Number:            req.Number,
		Type:              req.Type,
		Capacity:          req.Capacity,
		Description:       req.Description,
		IsAvailable:       req.IsAvailable,
		UpdatedAt:         time.Now(),
		MaintenanceStatus: req.MaintenanceStatus,
		UniqueTag:         req.UniqueTag,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return castServiceWorkplaceToProto(workplace), nil
}

func (s *serverAPI) DeleteWorkplace(ctx context.Context, req *proto_gen.DeleteWorkplaceRequest) (*proto_gen.DeleteWorkplaceResponse, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	err := s.workplaces.DeleteWorkplace(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.DeleteWorkplaceResponse{Success: true}, nil
}
func (s *serverAPI) CheckWorkplaceAvailability(ctx context.Context, req *proto_gen.CheckWorkplaceAvailabilityRequest) (*proto_gen.WorkplaceAvailabilityResponse, error) {
	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	available, timeSlots, err := s.workplaces.CheckWorkplaceAvailability(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	result := new(proto_gen.WorkplaceAvailabilityResponse)
	result.IsAvailable = available
	result.UnavailableSlots = make([]*proto_gen.TimeSlot, 0)
	for _, timeSlot := range timeSlots {
		result.UnavailableSlots = append(result.UnavailableSlots, &proto_gen.TimeSlot{
			StartTime: timestamppb.New(timeSlot.From),
			EndTime:   timestamppb.New(timeSlot.To),
			Reason:    timeSlot.Reason,
		})
	}
	return result, nil
}

func castServiceWorkplaceToProto(workplace services.Workplace) *proto_gen.Workplace {
	protoWorkplace := &proto_gen.Workplace{
		Id:                workplace.Id,
		Address:           workplace.Address,
		Zone:              workplace.Zone,
		Floor:             workplace.Floor,
		Number:            workplace.Number,
		Type:              workplace.Type,
		Capacity:          workplace.Capacity,
		Description:       workplace.Description,
		IsAvailable:       workplace.IsAvailable,
		CreatedAt:         timestamppb.New(workplace.CreatedAt),
		UpdatedAt:         timestamppb.New(workplace.UpdatedAt),
		MaintenanceStatus: workplace.MaintenanceStatus,
	}
	protoWorkplace.Items = make([]*proto_gen.Item, 0)
	for _, item := range workplace.Items {
		protoWorkplace.Items = append(protoWorkplace.Items, &proto_gen.Item{
			Id:          item.Id,
			Type:        item.Type,
			Name:        item.Name,
			Condition:   item.Condition,
			WorkplaceId: item.WorkplaceId,
			Quantity:    item.Quantity,
			IsAvailable: item.IsAvailable,
			CreatedAt:   timestamppb.New(item.CreatedAt),
			UpdatedAt:   timestamppb.New(item.UpdatedAt),
		})
	}
	return protoWorkplace
}
