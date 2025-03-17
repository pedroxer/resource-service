package storage

import database_models "github.com/pedroxer/resource-service/internal/models/database"

func (s Storage) GetWorkplaces(zone, floor, workplaceType string, capacity int64, isAvailable bool, page, pageSize int64) ([]database_models.Workplace, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) GetWorkplacesById(id int64) (database_models.Workplace, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) CreateWorkplace(workplace database_models.Workplace) (database_models.Workplace, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) UpdateWorkplace(workplace database_models.Workplace) (database_models.Workplace, error) {
	//TODO implement me
	panic("implement me")
}

func (s Storage) DeleteWorkplace(id int64) error {
	//TODO implement me
	panic("implement me")
}
