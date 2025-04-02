package storage

import (
	"context"
	"fmt"
	"github.com/pedroxer/resource-service/internal/models"
	"github.com/pedroxer/resource-service/internal/utills"
	"strings"
)

var parkingSpaceColumnsMap = map[string]SearchField{
	"id":           {NameWhere: "space.id", NameOrder: "space.id"},
	"address":      {NameWhere: "space.address", NameOrder: "space.address"},
	"zone":         {NameWhere: "space.zone", NameOrder: "space.zone"},
	"number":       {NameWhere: "space.number", NameOrder: "space.number"},
	"type":         {NameWhere: "space.type", NameOrder: "space.type"},
	"is_available": {NameWhere: "space.is_available", NameOrder: "space.is_available"},
	"created_at":   {NameWhere: "space.created_at", NameOrder: "space.created_at"},
	"updated_at":   {NameWhere: "space.updated_at", NameOrder: "space.updated_at"},
}

var lotsFields = []string{
	"id",
	"address",
	"zone",
	"number",
	"type",
	"is_available",
	"created_at",
	"updated_at",
}

func (s *Storage) GetParkingLots(ctx context.Context, filters []Field, page int64) ([]models.ParkingPlace, int64, error) {
	from := ` FROM resource_service.parking_spaces space`
	selectQuery := "SELECT " + strings.Join(lotsFields, ", ") + from

	countQuery := `SELECT count(*) FROM (` + selectQuery
	where, err := GenerateSearch(parkingSpaceColumnsMap, filters)
	if err != nil {
		s.logger.Warn(err.Error())
		return nil, 0, err
	}
	var conditions strings.Builder
	if len(filters) != 0 {
		conditions.WriteString(" WHERE")
		conditions.WriteString(where)
	}
	selectQuery += conditions.String() + GenerateLimits(page, utills.PageSize)

	rows, err := s.db.Query(ctx, selectQuery)
	if err != nil {
		s.logger.Warn(err.Error())
		return nil, 0, err
	}
	defer rows.Close()
	var lots []models.ParkingPlace
	var lotsCount int64
	for rows.Next() {
		var lot models.ParkingPlace
		if err := rows.Scan(
			&lot.Id,
			&lot.Address,
			&lot.Zone,
			&lot.Number,
			&lot.Type,
			&lot.IsAvailable,
			&lot.CreatedAt,
			&lot.UpdatedAt,
		); err != nil {
			s.logger.Warn("error getting spaces", err.Error())
			return nil, 0, err
		}
		lots = append(lots, lot)
	}

	countQuery += conditions.String()
	countQuery += `) as cnt`
	if err := s.db.QueryRow(ctx, countQuery).Scan(&lotsCount); err != nil {
		s.logger.Warn("error getting amount", err.Error())
		return nil, 0, err
	}

	return lots, lotsCount, nil
}

func (s *Storage) GetParkingLotById(ctx context.Context, id int64) (models.ParkingPlace, error) {
	from := ` FROM resource_service.parking_spaces space`
	selectQuery := "SELECT " + strings.Join(lotsFields, ", ") + from
	selectQuery += " WHERE space.id = $1"
	var lot models.ParkingPlace
	if err := s.db.QueryRow(ctx, selectQuery, id).Scan(
		&lot.Id,
		&lot.Address,
		&lot.Zone,
		&lot.Number,
		&lot.Type,
		&lot.IsAvailable,
		&lot.CreatedAt,
		&lot.UpdatedAt,
	); err != nil {
		s.logger.Warn("error getting space", err.Error())
		return models.ParkingPlace{}, err
	}

	return lot, nil
}

func (s *Storage) CreateParkingLot(ctx context.Context, parkingSpace models.ParkingPlace) (int64, error) {
	query := `INSERT INTO resource_service.parking_spaces (address, zone, number, type, is_available) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int64
	if err := s.db.QueryRow(ctx, query, parkingSpace.Address, parkingSpace.Zone, parkingSpace.Number, parkingSpace.Type, parkingSpace.IsAvailable).Scan(&id); err != nil {
		s.logger.Warn("error creating space", err.Error())
		return 0, err
	}
	return id, nil
}

func (s *Storage) UpdateParkingLot(ctx context.Context, id int64, updateFields []Field) (models.ParkingPlace, error) {
	updateQuery := `UPDATE resource_service.parking_spaces SET `
	updateColumns, err := GenerateUpdates(parkingSpaceColumnsMap, updateFields)
	if err != nil {
		s.logger.Warn("cannot generate columns", err.Error())
		return models.ParkingPlace{}, err
	}
	if len(updateColumns) == 0 {
		s.logger.Warn("not enough data to update")
		return models.ParkingPlace{}, fmt.Errorf("not enough data to update")
	}
	updateQuery += updateColumns + fmt.Sprintf(" WHERE id = %d RETURNING *", id)
	var result models.ParkingPlace
	if err := s.db.QueryRow(ctx, updateQuery).Scan(
		&result.Id,
		&result.Address,
		&result.Zone,
		&result.Number,
		&result.Type,
		&result.IsAvailable,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		s.logger.Warn("cannot update space", err.Error())
		return models.ParkingPlace{}, err
	}

	return result, nil
}

func (s *Storage) DeleteParkingLot(ctx context.Context, id int64) error {
	query := `DELETE FROM resource_service.parking_spaces WHERE id = $1`
	if _, err := s.db.Exec(ctx, query, id); err != nil {
		s.logger.Warn("cannot delete space", err.Error())
		return err
	}
	return nil
}
