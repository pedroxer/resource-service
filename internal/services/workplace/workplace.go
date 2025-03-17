package workplace

import (
	"github.com/pedroxer/resource-service/internal/models/database"
	"github.com/pedroxer/resource-service/internal/models/services"
)

type WorkplaceGetter interface {
	GetWorkplaces(zone, floor, workplaceType string, capacity int64, isAvailable bool, page, pageSize int64) ([]database_models.Workplace, int64, error)
	GetWorkplacesById(id int64) (database_models.Workplace, error)
}
type WorkplaceCreater interface {
	CreateWorkplace(workplace database_models.Workplace) (database_models.Workplace, error)
	UpdateWorkplace(workplace database_models.Workplace) (database_models.Workplace, error)
	DeleteWorkplace(id int64) error
}
type DefaultWorkplaceService struct {
}

func (d DefaultWorkplaceService) GetWorkplaces(zone, floor, workplaceType string, capacity int64, isAvailable bool, page, pageSize int64) ([]services.Workplace, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultWorkplaceService) GetWorkplacesById(id int64) (services.Workplace, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultWorkplaceService) CreateWorkplace(workplace services.Workplace) (services.Workplace, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultWorkplaceService) UpdateWorkplace(workplace services.Workplace) (services.Workplace, error) {
	//TODO implement me
	panic("implement me")
}

func (d DefaultWorkplaceService) DeleteWorkplace(id int64) error {
	//TODO implement me
	panic("implement me")
}

func (d DefaultWorkplaceService) CheckWorkplaceAvailability(id int64) (bool, []services.TimeSlot, error) {
	//TODO implement me
	panic("implement me")
}
