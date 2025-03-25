package workplace

import (
	"context"
	"github.com/pedroxer/resource-service/internal/models"
	"github.com/pedroxer/resource-service/internal/storage"
	"github.com/pedroxer/resource-service/internal/utills"
	log "github.com/sirupsen/logrus"
)

type WorkplaceGetter interface {
	GetWorkplaces(ctx context.Context, filters []storage.Field, page int64) ([]models.Workplace, int64, error)
	GetWorkplacesById(ctx context.Context, id int64) (models.Workplace, error)
	GetWorkplaceByUniqueTag(ctx context.Context, uniqueTag string) (models.Workplace, error)
}
type WorkplaceCreater interface {
	CreateWorkplace(ctx context.Context, workplace models.Workplace) (models.Workplace, error)
	UpdateWorkplace(ctx context.Context, id int64, updateFields []storage.Field) (models.Workplace, error)
	DeleteWorkplace(ctx context.Context, id int64) error
}
type DefaultWorkplaceService struct {
	logger  *log.Logger
	creater WorkplaceCreater
	getter  WorkplaceGetter
}

func NewDefaultWorkplaceService(storage *storage.Storage, logger *log.Logger) *DefaultWorkplaceService {
	return &DefaultWorkplaceService{
		logger:  logger,
		creater: storage,
		getter:  storage,
	}
}

func (d DefaultWorkplaceService) GetWorkplaces(ctx context.Context, zone, workplaceType string, floor int64, capacity int64, isAvailable bool, page, pageSize int64) ([]models.Workplace, int64, error) {
	filters := make([]storage.Field, 0)
	if zone != "" {
		filters = append(filters, storage.Field{
			Name:  "zone",
			Value: zone,
		})
	}
	if floor != 0 {
		filters = append(filters, storage.Field{
			Name:  "floor",
			Value: floor,
		})
	}
	if workplaceType != "" {
		filters = append(filters, storage.Field{
			Name:  "type",
			Value: workplaceType,
		})
	}
	if capacity != 0 {
		filters = append(filters, storage.Field{
			Name:  "capacity",
			Value: capacity,
		})
	}

	filters = append(filters, storage.Field{
		Name:  "is_available",
		Value: isAvailable,
	})

	workplaces, amount, err := d.getter.GetWorkplaces(ctx, filters, page)
	if err != nil {
		d.logger.Warn(err.Error())
		return nil, 0, err
	}
	return workplaces, amount, nil
}

func (d DefaultWorkplaceService) GetWorkplacesById(ctx context.Context, id int64) (models.Workplace, error) {
	workplace, err := d.getter.GetWorkplacesById(ctx, id)
	if err != nil {
		d.logger.Warn(err.Error())
		return models.Workplace{}, err
	}
	return workplace, err
}

func (d DefaultWorkplaceService) GetWorkplaceByUniqueTag(ctx context.Context, uniqueTag string) (models.Workplace, error) {
	workplace, err := d.getter.GetWorkplaceByUniqueTag(ctx, uniqueTag)
	if err != nil {
		d.logger.Warn(err.Error())
		return models.Workplace{}, err
	}
	return workplace, err
}
func (d DefaultWorkplaceService) CreateWorkplace(ctx context.Context, workplace models.Workplace) (models.Workplace, error) {
	uniqueTag := utills.GenerateUniqueTag(workplace.Address, workplace.Zone, workplace.Type, workplace.Floor, workplace.Number)
	workplace.UniqueTag = uniqueTag
	finalWorkplace, err := d.creater.CreateWorkplace(ctx, workplace)
	if err != nil {
		d.logger.Warn(err.Error())
		return models.Workplace{}, err
	}
	return finalWorkplace, nil
}

func (d DefaultWorkplaceService) UpdateWorkplace(ctx context.Context, workplace models.Workplace) (models.Workplace, error) {
	updateFields := make([]storage.Field, 0)
	if workplace.Address != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "address",
			Value: workplace.Address,
		})
	}
	if workplace.Zone != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "zone",
			Value: workplace.Zone,
		})
	}
	if workplace.Floor != 0 {
		updateFields = append(updateFields, storage.Field{
			Name:  "floor",
			Value: workplace.Floor,
		})
	}
	if workplace.Number != 0 {
		updateFields = append(updateFields, storage.Field{
			Name:  "number",
			Value: workplace.Number,
		})
	}
	if workplace.Type != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "type",
			Value: workplace.Type,
		})
	}
	if workplace.Capacity != 0 {
		updateFields = append(updateFields, storage.Field{
			Name:  "capacity",
			Value: workplace.Capacity,
		})
	}
	if workplace.Description != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "description",
			Value: workplace.Description,
		})
	}
	if workplace.IsAvailable != false {
		updateFields = append(updateFields, storage.Field{
			Name:  "is_available",
			Value: workplace.IsAvailable,
		})
	}
	if workplace.MaintenanceStatus != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "maintenance_status",
			Value: workplace.MaintenanceStatus,
		})
	}
	if workplace.UniqueTag != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "unique_tag",
			Value: workplace.UniqueTag,
		})
	}

	finalWorkplace, err := d.creater.UpdateWorkplace(ctx, workplace.Id, updateFields)
	if err != nil {
		d.logger.Warn(err.Error())
		return models.Workplace{}, err
	}
	return finalWorkplace, nil
}

func (d DefaultWorkplaceService) DeleteWorkplace(ctx context.Context, id int64) error {
	err := d.creater.DeleteWorkplace(ctx, id)
	if err != nil {
		d.logger.Warn(err.Error())
		return err
	}
	return nil
}
