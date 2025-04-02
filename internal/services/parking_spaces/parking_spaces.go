package parking_spaces

import (
	"context"
	"github.com/pedroxer/resource-service/internal/models"
	"github.com/pedroxer/resource-service/internal/storage"
	log "github.com/sirupsen/logrus"
)

type ParkingSpaceGetter interface {
	GetParkingLots(ctx context.Context, filters []storage.Field, page int64) ([]models.ParkingPlace, int64, error)
	GetParkingLotById(ctx context.Context, id int64) (models.ParkingPlace, error)
}

type ParkingSpaceCreator interface {
	CreateParkingLot(ctx context.Context, parkingSpace models.ParkingPlace) (int64, error)
	UpdateParkingLot(ctx context.Context, id int64, updateFields []storage.Field) (models.ParkingPlace, error)
	DeleteParkingLot(ctx context.Context, id int64) error
}
type DefaultParkingSpaceService struct {
	logger  *log.Logger
	creater ParkingSpaceCreator
	getter  ParkingSpaceGetter
}

func NewDefaultParkingSpaceService(storage *storage.Storage, logger *log.Logger) *DefaultParkingSpaceService {
	return &DefaultParkingSpaceService{
		logger:  logger,
		creater: storage,
		getter:  storage,
	}
}

func (d DefaultParkingSpaceService) GetParkingSpaces(ctx context.Context, address, zone, spaceType string, isAvailable bool, page int64) ([]models.ParkingPlace, int64, error) {
	filter := make([]storage.Field, 0)
	if address != "" {
		filter = append(filter, storage.Field{
			Name:  "address",
			Value: address,
		})
	}
	if zone != "" {
		filter = append(filter, storage.Field{
			Name:  "zone",
			Value: zone,
		})
	}
	if spaceType != "" {
		filter = append(filter, storage.Field{
			Name:  "type",
			Value: spaceType,
		})
	}
	if isAvailable {
		filter = append(filter, storage.Field{
			Name:  "is_available",
			Value: isAvailable,
		})
	}
	spaces, amount, err := d.getter.GetParkingLots(ctx, filter, page)
	if err != nil {
		d.logger.Warn("error getting spaces", err.Error())
		return nil, 0, err
	}
	return spaces, amount, nil
}

func (d DefaultParkingSpaceService) GetParkingSpaceById(ctx context.Context, id int64) (models.ParkingPlace, error) {
	lot, err := d.getter.GetParkingLotById(ctx, id)
	if err != nil {
		d.logger.Warn("error getting space", err.Error())
		return models.ParkingPlace{}, err
	}
	return lot, nil
}

func (d DefaultParkingSpaceService) CreateParkingSpace(ctx context.Context, parkingSpace models.ParkingPlace) (int64, error) {
	id, err := d.creater.CreateParkingLot(ctx, parkingSpace)
	if err != nil {
		d.logger.Warn("error creating space", err.Error())
		return 0, err
	}
	return id, nil
}

func (d DefaultParkingSpaceService) UpdateParkingSpace(ctx context.Context, parkingSpace models.ParkingPlace) (models.ParkingPlace, error) {
	updateFields := make([]storage.Field, 0)
	if parkingSpace.Address != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "address",
			Value: parkingSpace.Address,
		})
	}
	if parkingSpace.Zone != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "zone",
			Value: parkingSpace.Zone,
		})
	}
	if parkingSpace.Type != "" {
		updateFields = append(updateFields, storage.Field{
			Name:  "type",
			Value: parkingSpace.Type,
		})
	}
	if parkingSpace.IsAvailable {
		updateFields = append(updateFields, storage.Field{
			Name:  "is_available",
			Value: parkingSpace.IsAvailable,
		})
	}

	parkingSpace, err := d.creater.UpdateParkingLot(ctx, parkingSpace.Id, updateFields)
	if err != nil {
		d.logger.Warn("error updating space", err.Error())
		return models.ParkingPlace{}, err
	}
	return parkingSpace, nil
}

func (d DefaultParkingSpaceService) DeleteParkingSpace(ctx context.Context, id int64) error {
	err := d.creater.DeleteParkingLot(ctx, id)
	if err != nil {
		d.logger.Warn("error deleting space", err.Error())
		return err
	}
	return nil
}
