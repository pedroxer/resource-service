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

type WorkplaceService interface {
	GetWorkplaces(ctx context.Context, zone, workplaceType string, floor, capacity int64, isAvailable, withItems bool, page, pageSize int64) ([]models.Workplace, int64, error)
	GetWorkplacesById(ctx context.Context, id int64) (models.Workplace, error)
	GetWorkplaceByUniqueTag(ctx context.Context, uniqueTag string) (models.Workplace, error)
	CreateWorkplace(ctx context.Context, workplace models.Workplace) (models.Workplace, error)
	UpdateWorkplace(ctx context.Context, workplace models.Workplace) (models.Workplace, error)
	DeleteWorkplace(ctx context.Context, id int64) error
}

func (s *serverAPI) GetWorkplaces(ctx context.Context, req *proto_gen.GetWorkplacesRequest) (*proto_gen.GetWorkplacesResponse, error) {
	if req.GetPage() == 0 {
		req.Page = 1
	}
	workplaces, amount, err := s.workplaces.GetWorkplaces(ctx, req.Zone, req.Type, req.Floor, req.Capacity, req.IsAvailable, req.WithItems, req.Page, utills.PageSize)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := new(proto_gen.GetWorkplacesResponse)
	response.Workplaces = make([]*proto_gen.Workplace, 0)
	for _, workplace := range workplaces {
		response.Workplaces = append(response.Workplaces, castServiceWorkplaceToProto(workplace))
	}
	response.PageSize = utills.PageSize
	response.TotalCount = amount/utills.PageSize + 1
	response.Page = req.Page
	return response, nil
}

func (s *serverAPI) GetWorkplaceById(ctx context.Context, req *proto_gen.GetWorkplaceByIdRequest) (*proto_gen.Workplace, error) {
	workplace, err := s.workplaces.GetWorkplacesById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return castServiceWorkplaceToProto(workplace), nil
}

func (s *serverAPI) GetWorkplaceByUniqueTag(ctx context.Context, req *proto_gen.GetWorkplaceByUniqueTagRequest) (*proto_gen.Workplace, error) {
	if req.GetUniqueTag() == "" {
		return nil, status.Error(codes.InvalidArgument, "unique tag is required")
	}
	workplace, err := s.workplaces.GetWorkplaceByUniqueTag(ctx, req.UniqueTag)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return castServiceWorkplaceToProto(workplace), nil
}
func (s *serverAPI) CreateWorkplace(ctx context.Context, req *proto_gen.CreateWorkplaceRequest) (*proto_gen.Workplace, error) {
	if req.GetAddress() == "" || req.GetType() == "" || req.Capacity == 0 || req.Number == 0 || req.Floor == 0 {
		return nil, status.Error(codes.InvalidArgument, "address, type, capacity, number and floor are required")
	}
	workplace, err := s.workplaces.CreateWorkplace(ctx, models.Workplace{
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
	workplace, err := s.workplaces.UpdateWorkplace(ctx, models.Workplace{
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
	err := s.workplaces.DeleteWorkplace(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.DeleteWorkplaceResponse{Success: true}, nil
}

func castServiceWorkplaceToProto(workplace models.Workplace) *proto_gen.Workplace {
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
			CreatedAt:   timestamppb.New(item.CreatedAt),
			UpdatedAt:   timestamppb.New(item.UpdatedAt),
		})
	}
	return protoWorkplace
}
