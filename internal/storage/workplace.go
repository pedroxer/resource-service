package storage

import (
	"context"
	"fmt"
	"github.com/pedroxer/resource-service/internal/models"
	"github.com/pedroxer/resource-service/internal/utills"
	"strings"
	"time"
)

var wrokplaceColumnsMap = map[string]SearchField{
	"id":                  {NameWhere: "workplace.id", NameOrder: "workplace.id"},
	"address":             {NameWhere: "workplace.address", NameOrder: "workplace.address"},
	"zone":                {NameWhere: "workplace.zone", NameOrder: "workplace.zone"},
	"floor":               {NameWhere: "workplace.floor", NameOrder: "workplace.floor"},
	"number":              {NameWhere: "workplace.number", NameOrder: "workplace.number"},
	"type":                {NameWhere: "workplace.type", NameOrder: "workplace.type"},
	"capacity":            {NameWhere: "workplace.capacity", NameOrder: "workplace.capacity"},
	"description":         {NameWhere: "workplace.description", NameOrder: "workplace.description"},
	"is_available":        {NameWhere: "workplace.is_available", NameOrder: "workplace.is_available"},
	"maintainance_status": {NameWhere: "workplace.maintainance_status", NameOrder: "workplace.maintainance_status"},
	"created_at":          {NameWhere: "workplace.created_at", NameOrder: "workplace.created_at"},
	"updated_at":          {NameWhere: "workplace.updated_at", NameOrder: "workplace.updated_at"},
	"unique_tag":          {NameWhere: "workplace.unique_tag", NameOrder: "workplace.unique_tag"},
}
var workplacesFileds = []string{
	"workplace.id",
	"workplace.address",
	"workplace.zone",
	"workplace.floor",
	"workplace.number",
	"workplace.type",
	"workplace.capacity",
	"workplace.description",
	"workplace.is_available",
	"workplace.maintainance_status",
	"workplace.created_at",
	"workplace.updated_at",
	"workplace.unique_tag",
}

func (s *Storage) GetWorkplaces(ctx context.Context, filters []Field, withItems bool, page int64) ([]models.Workplace, int64, error) {

	from := ` FROM resource_service.workplace`
	selectQuery := "SELECT " + strings.Join(workplacesFileds, ", ") + from

	countQuery := `SELECT count(*) FROM (` + selectQuery
	where, err := GenerateSearch(wrokplaceColumnsMap, filters)
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
	var workplaces []models.Workplace
	var workplacesCount int64
	var workplaceIds []int64
	for rows.Next() {
		var workplace models.Workplace

		if err := rows.Scan(
			&workplace.Id,
			&workplace.Address,
			&workplace.Zone,
			&workplace.Floor,
			&workplace.Number,
			&workplace.Type,
			&workplace.Capacity,
			&workplace.Description,
			&workplace.IsAvailable,
			&workplace.MaintenanceStatus,
			&workplace.CreatedAt,
			&workplace.UpdatedAt,
			&workplace.UniqueTag,
		); err != nil {
			return nil, 0, err

		}
		workplaceIds = append(workplaceIds, workplace.Id)
		workplaces = append(workplaces, workplace)
	}
	countQuery += conditions.String()
	countQuery += ") as cnt"
	if withItems {
		items, err := s.GetItemsByWorkplaceIds(ctx, workplaceIds)
		if err != nil {
			s.logger.Warn(err.Error())
			return nil, 0, err
		}
		for i, workplace := range workplaces {
			workplaces[i].Items = items[workplace.Id]
		}
	}
	if err := s.db.QueryRow(ctx, countQuery).Scan(&workplacesCount); err != nil {
		s.logger.Warn(err.Error())
		return nil, 0, err
	}
	return workplaces, workplacesCount, nil
}

func (s *Storage) GetWorkplacesById(ctx context.Context, id int64) (models.Workplace, error) {
	from := ` FROM resource_service.workplace`
	selectQuery := "SELECT " + strings.Join(workplacesFileds, ", ") + from
	selectQuery += " WHERE workplace.id = $1"
	var workplace models.Workplace
	if err := s.db.QueryRow(ctx, selectQuery, id).Scan(
		&workplace.Id,
		&workplace.Address,
		&workplace.Zone,
		&workplace.Floor,
		&workplace.Number,
		&workplace.Type,
		&workplace.Capacity,
		&workplace.Description,
		&workplace.IsAvailable,
		&workplace.MaintenanceStatus,
		&workplace.CreatedAt,
		&workplace.UpdatedAt,
		&workplace.UniqueTag,
	); err != nil {
		s.logger.Warn(err.Error())
		return models.Workplace{}, err
	}
	items, err := s.GetItemsByWorkplaceId(ctx, workplace.Id)
	if err != nil {
		s.logger.Warn(err.Error())
		return models.Workplace{}, err
	}
	workplace.Items = items
	return workplace, nil
}

func (s *Storage) GetWorkplaceByUniqueTag(ctx context.Context, uniqueTag string) (models.Workplace, error) {
	from := ` FROM resource_service.workplace`
	selectQuery := "SELECT " + strings.Join(workplacesFileds, ", ") + from
	selectQuery += " WHERE workplace.unique_tag = $1"
	var workplace models.Workplace
	if err := s.db.QueryRow(ctx, selectQuery, uniqueTag).Scan(
		&workplace.Id,
		&workplace.Address,
		&workplace.Zone,
		&workplace.Floor,
		&workplace.Number,
		&workplace.Type,
		&workplace.Capacity,
		&workplace.Description,
		&workplace.IsAvailable,
		&workplace.MaintenanceStatus,
		&workplace.CreatedAt,
		&workplace.UpdatedAt,
		&workplace.UniqueTag,
	); err != nil {
		s.logger.Warn(err.Error())
		return models.Workplace{}, err
	}
	items, err := s.GetItemsByWorkplaceId(ctx, workplace.Id)
	if err != nil {
		s.logger.Warn(err.Error())
		return models.Workplace{}, err
	}
	workplace.Items = items
	return workplace, nil
}

func (s *Storage) CreateWorkplace(ctx context.Context, workplace models.Workplace) (models.Workplace, error) {
	query := `INSERT INTO resource_service."workplace" (
		address, zone, floor, number, type, capacity, description, is_available, maintainance_status, unique_tag
	) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;`

	var id int64
	if err := s.db.QueryRow(ctx, query,
		workplace.Address,
		workplace.Zone,
		workplace.Floor,
		workplace.Number,
		workplace.Type,
		workplace.Capacity,
		workplace.Description,
		workplace.IsAvailable,
		workplace.MaintenanceStatus,
		workplace.UniqueTag,
	).Scan(&id); err != nil {
		s.logger.Warn("failed to create workplace", err.Error())
		return models.Workplace{}, err
	}
	return models.Workplace{
		Id:                id,
		Address:           workplace.Address,
		Zone:              workplace.Zone,
		Floor:             workplace.Floor,
		Number:            workplace.Number,
		Type:              workplace.Type,
		Capacity:          workplace.Capacity,
		Description:       workplace.Description,
		IsAvailable:       workplace.IsAvailable,
		CreatedAt:         workplace.CreatedAt,
		UpdatedAt:         time.Time{},
		MaintenanceStatus: workplace.MaintenanceStatus,
		UniqueTag:         workplace.UniqueTag,
		Items:             nil,
	}, nil
}

func (s *Storage) UpdateWorkplace(ctx context.Context, id int64, updateFields []Field) (models.Workplace, error) {
	var updateQueryBuilder strings.Builder
	updateQueryBuilder.WriteString("UPDATE resource_service.workplace SET ")
	for _, field := range updateFields {
		updateQueryBuilder.WriteString(fmt.Sprintf("%s = '%s' AND ", field.Name, fmt.Sprint(field.Value)))
	}
	updateQuery := updateQueryBuilder.String()
	updateQuery = strings.Trim(updateQuery, "AND ")
	updateQuery += fmt.Sprintf(" WHERE id = %d RETURNING *", id)

	var result models.Workplace
	if err := s.db.QueryRow(ctx, updateQuery).Scan(
		&result.Id,
		&result.Address,
		&result.Zone,
		&result.Floor,
		&result.Number,
		&result.Type,
		&result.Capacity,
		&result.Description,
		&result.IsAvailable,
		&result.MaintenanceStatus,
		&result.UniqueTag,
		&result.CreatedAt,
		&result.UpdatedAt,
	); err != nil {
		s.logger.Warn(err.Error())
		return models.Workplace{}, err
	}
	return result, nil
}

func (s *Storage) DeleteWorkplace(ctx context.Context, id int64) error {
	query := `DELETE FROM resource_service.workplace WHERE id = $1`

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		s.logger.Warn(err.Error())
		return err
	}
	return nil
}
