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

type ItemService interface {
	GetItems(itemType, name string, conditionId, workplaceId int64, isAvailable bool, page, pageSize int64) ([]models.Item, int64, error)
	GetItemById(id int64) (models.Item, error)
	CreateItem(item *models.Item) (int64, error)
	UpdateItem(item *models.Item) (models.Item, error)
	DeleteItem(id int64) error
}

func (s *serverAPI) GetItems(ctx context.Context, req *proto_gen.GetItemsRequest) (*proto_gen.GetItemsResponse, error) {
	if req.GetPage() == 0 {
		req.Page = 1
	}
	items, amount, err := s.items.GetItems(req.Type, req.Name, req.ConditionId, req.WorkplaceId, req.IsAvailable, req.Page, utills.PageSize)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := new(proto_gen.GetItemsResponse)
	response.Items = make([]*proto_gen.Item, 0)
	for _, item := range items {
		response.Items = append(response.Items, &proto_gen.Item{
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
	response.PageSize = utills.PageSize
	response.TotalCount = amount / utills.PageSize
	response.Page = req.Page
	return response, nil
}

func (s *serverAPI) GetItemById(ctx context.Context, req *proto_gen.GetItemByIdRequest) (*proto_gen.Item, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	item, err := s.items.GetItemById(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.Item{
		Id:          item.Id,
		Type:        item.Type,
		Name:        item.Name,
		Condition:   item.Condition,
		WorkplaceId: item.WorkplaceId,
		Quantity:    item.Quantity,
		IsAvailable: item.IsAvailable}, nil
}

func (s *serverAPI) CreateItem(ctx context.Context, req *proto_gen.CreateItemRequest) (*proto_gen.Item, error) {
	item := models.Item{
		Type:        req.Type,
		Name:        req.Name,
		Condition:   utills.IdsToConditions[req.ConditionId],
		WorkplaceId: req.WorkplaceId,
		Quantity:    req.Quantity,
		IsAvailable: req.IsAvailable,
	}
	id, err := s.items.CreateItem(&item)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.Item{
		Id:          id,
		Type:        item.Type,
		Name:        item.Name,
		Condition:   item.Condition,
		WorkplaceId: item.WorkplaceId,
		Quantity:    item.Quantity,
		IsAvailable: item.IsAvailable,
		CreatedAt:   timestamppb.New(time.Now()),
		UpdatedAt:   timestamppb.New(time.Now()),
	}, nil
}

func (s *serverAPI) UpdateItem(ctx context.Context, req *proto_gen.UpdateItemRequest) (*proto_gen.Item, error) {
	item := models.Item{
		Id:          req.Id,
		Type:        req.Type,
		Name:        req.Name,
		Condition:   utills.IdsToConditions[req.ConditionId],
		WorkplaceId: req.WorkplaceId,
		Quantity:    req.Quantity,
		IsAvailable: req.IsAvailable,
	}
	resp, err := s.items.UpdateItem(&item)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.Item{
		Id:          resp.Id,
		Type:        resp.Type,
		Name:        resp.Name,
		Condition:   resp.Condition,
		WorkplaceId: resp.WorkplaceId,
		Quantity:    resp.Quantity,
		IsAvailable: resp.IsAvailable,
		CreatedAt:   timestamppb.New(resp.CreatedAt),
		UpdatedAt:   timestamppb.New(resp.UpdatedAt),
	}, nil
}

func (s *serverAPI) DeleteItem(ctx context.Context, req *proto_gen.DeleteItemRequest) (*proto_gen.DeleteItemResponse, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	err := s.items.DeleteItem(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto_gen.DeleteItemResponse{Success: true}, nil
}
