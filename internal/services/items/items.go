package items

import (
	"context"
	"github.com/pedroxer/resource-service/internal/models"
	"github.com/pedroxer/resource-service/internal/storage"
	log "github.com/sirupsen/logrus"
)

type ItemGetter interface {
	GetItemsByWorkplace(ctx context.Context, workplaceId int64, page, pageSize int64) ([]models.Item, int64, error)
	GetItems(ctx context.Context, filters []storage.Field, page int64) ([]models.Item, int64, error)
	GetItemById(ctx context.Context, id int64) (models.Item, error)
}

type ItemCreater interface {
	CreateItem(ctx context.Context, item models.Item) (int64, error)
	UpdateItem(ctx context.Context, id int64, updateFields []storage.Field) (models.Item, error)
	DeleteItem(ctx context.Context, id int64) error
}
type DefaultItemService struct {
	logger  *log.Logger
	creater ItemCreater
	getter  ItemGetter
}

func NewDefaultItemService(storage *storage.Storage, logger *log.Logger) *DefaultItemService {
	return &DefaultItemService{
		logger:  logger,
		creater: storage,
		getter:  storage,
	}
}

func (d DefaultItemService) GetItems(ctx context.Context, itemType, name string, conditionId, workplaceId int64, page, pageSize int64) ([]models.Item, int64, error) {
	filters := make([]storage.Field, 0)
	if itemType != "" {
		filters = append(filters, storage.Field{
			Name:  "type",
			Value: itemType,
		})
	}
	if name != "" {
		filters = append(filters, storage.Field{
			Name:  "name",
			Value: name,
		})
	}
	if conditionId != 0 {
		filters = append(filters, storage.Field{
			Name:  "condition_id",
			Value: conditionId,
		})
	}
	if workplaceId != 0 {
		filters = append(filters, storage.Field{
			Name:  "workplace_id",
			Value: workplaceId,
		})
	}
	items, amount, err := d.getter.GetItems(ctx, filters, page)
	if err != nil {
		d.logger.Warn("error getting items", err.Error())
		return nil, 0, err
	}

	return items, amount, nil
}

func (d DefaultItemService) GetItemById(ctx context.Context, id int64) (models.Item, error) {
	item, err := d.getter.GetItemById(ctx, id)
	if err != nil {
		d.logger.Warn("error getting item", err.Error())
		return models.Item{}, err
	}
	return item, nil
}

func (d DefaultItemService) CreateItem(ctx context.Context, item models.Item) (int64, error) {
	id, err := d.creater.CreateItem(ctx, item)
	if err != nil {
		d.logger.Warn("error creating item", err.Error())
		return 0, err
	}
	return id, nil
}

func (d DefaultItemService) UpdateItem(ctx context.Context, item models.Item) (models.Item, error) {
	updateFields := make([]storage.Field, 0)
	if item.Type != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "type",
			Value: item.Type,
		})
	}
	if item.Name != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "name",
			Value: item.Name,
		})
	}
	if item.Condition != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "condition",
			Value: item.Condition,
		})
	}
	if item.WorkplaceId != 0 {
		updateFields = append(updateFields, storage.Field{
			Name:  "workplace_id",
			Value: item.WorkplaceId,
		})
	}
	item, err := d.creater.UpdateItem(ctx, item.Id, updateFields)
	if err != nil {
		d.logger.Warn("error updating item", err.Error())
		return models.Item{}, err
	}
	return item, nil
}

func (d DefaultItemService) DeleteItem(ctx context.Context, id int64) error {
	err := d.creater.DeleteItem(ctx, id)
	if err != nil {
		d.logger.Warn("error deleting item", err.Error())
		return err
	}
	return nil
}
