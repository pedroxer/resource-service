package items

import (
	"context"
	"github.com/pedroxer/resource-service/internal/models"
)

type ItemGetter interface {
	GetItemsByWorkplace(ctx context.Context, workplaceId int64, page, pageSize int64) ([]models.Item, int64, error)
}
type DefaultItemService struct {
}

func (d DefaultItemService) GetItems(itemType, name string, conditionId, workplaceId int64, page, pageSize int64) ([]models.Item, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultItemService) GetItemById(id int64) (models.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultItemService) CreateItem(item *models.Item) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultItemService) UpdateItem(item *models.Item) (models.Item, error) {
	//TODO implement me
	panic("implement me")

}

func (d DefaultItemService) DeleteItem(id int64) error {
	//TODO implement me
	panic("implement me")
}
