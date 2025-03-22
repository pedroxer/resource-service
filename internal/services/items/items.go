package items

import "github.com/pedroxer/resource-service/internal/models/services"

type DefaultItemService struct {
}

func (d DefaultItemService) GetItems(itemType, name string, conditionId, workplaceId int64, isAvailable bool, page, pageSize int64) ([]services.Item, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultItemService) GetItemById(id int64) (services.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultItemService) CreateItem(item *services.Item) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultItemService) UpdateItem(item *services.Item) (services.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultItemService) DeleteItem(id int64) error {
	//TODO implement me
	panic("implement me")
}
