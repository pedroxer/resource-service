package storage

import (
	"context"
	"github.com/pedroxer/resource-service/internal/models"
)

func (s *Storage) GetItemsByWorkplaceId(ctx context.Context, workplaceId int64) ([]models.Item, error) {
	query := `SELECT i.id,
	   i.name, 
	   i.type,
	   ic.value,
	   i.updated_at,
	   i.created_at FROM resource_service.items i 
		   
	INNER JOIN resource_service.item_conditions ic on ic.id = i.condition_id WHERE i.workplace_id = $1`

	var items []models.Item
	rows, err := s.db.Query(ctx, query, workplaceId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.Id, &item.Name, &item.Type, &item.Condition, &item.UpdatedAt, &item.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (s *Storage) GetItemsByWorkplaceIds(ctx context.Context, workplaceIds []int64) (map[int64][]models.Item, error) {
	query := `SELECT i.id,
	   i.name, 
	   i.type,
	   ic.value,
	   i.workplace_id,
	i.updated_at,
	i.created_at FROM resource_service.items i 
		   
	INNER JOIN resource_service.item_conditions ic on ic.id = i.condition_id WHERE i.workplace_id = ANY($1)`
	rows, err := s.db.Query(ctx, query, workplaceIds)
	if err != nil {
		s.logger.Warn(err.Error())
		return nil, err
	}
	var result = make(map[int64][]models.Item)
	for rows.Next() {
		var item models.Item
		err = rows.Scan(&item.Id, &item.Name, &item.Type, &item.Condition, &item.WorkplaceId, &item.UpdatedAt, &item.CreatedAt)
		if err != nil {
			s.logger.Warn(err.Error())
			return nil, err
		}
		if _, ok := result[item.WorkplaceId]; ok {
			result[item.WorkplaceId] = append(result[item.WorkplaceId], item)
		} else {
			result[item.WorkplaceId] = make([]models.Item, 0)
			result[item.WorkplaceId] = append(result[item.WorkplaceId], item)
		}
	}
	return result, nil
}
