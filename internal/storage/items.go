package storage

import (
	"context"
	"fmt"
	"github.com/pedroxer/resource-service/internal/models"
	"github.com/pedroxer/resource-service/internal/utills"
	"strings"
)

var itemsColumnsMap = map[string]SearchField{
	"id":           {NameWhere: "item.id", NameOrder: "item.id"},
	"type":         {NameWhere: "item.type", NameOrder: "item.type"},
	"name":         {NameWhere: "item.name", NameOrder: "item.name"},
	"condition_id": {NameWhere: "item.condition_id", NameOrder: "item.condition_id"},
	"workplace_id": {NameWhere: "item.workplace_id", NameOrder: "item.workplace_id"},
	"updated_at":   {NameWhere: "item.updated_at", NameOrder: "item.updated_at"},
	"created_at":   {NameWhere: "item.created_at", NameOrder: "item.created_at"},
}

var itemsFields = []string{
	"id",
	"type",
	"name",
	"condition_id",
	"workplace_id",
	"updated_at",
	"created_at",
}

func (s *Storage) GetItems(ctx context.Context, filters []Field, page int64) ([]models.Item, int64, error) {
	from := ` FROM resource_service.items item`
	selectQuery := "SELECT " + strings.Join(itemsFields, ", ") + from

	countQuery := `SELECT count(*) FROM (` + selectQuery
	where, err := GenerateSearch(itemsColumnsMap, filters)
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
	var items []models.Item
	var itemsCount int64
	for rows.Next() {
		var item models.Item
		var conditionId int64
		if err := rows.Scan(
			&item.Id,
			&item.Type,
			&item.Name,
			&conditionId,
			&item.WorkplaceId,
			&item.UpdatedAt,
			&item.CreatedAt,
		); err != nil {
			s.logger.Warn(err.Error())
			return nil, 0, err
		}
		condition := utills.IdsToConditions[conditionId]
		item.Condition = condition
		items = append(items, item)
	}
	countQuery += conditions.String()
	countQuery += `) as cnt`

	if err := s.db.QueryRow(ctx, countQuery).Scan(&itemsCount); err != nil {
		s.logger.Warn(err.Error())
		return nil, 0, err
	}

	return items, itemsCount, nil
}

func (s *Storage) GetItemById(ctx context.Context, id int64) (models.Item, error) {
	from := ` FROM resource_service.items item`
	selectQuery := "SELECT " + strings.Join(itemsFields, ", ") + from
	selectQuery += " WHERE item.id = $1"
	var (
		item        models.Item
		conditionId int64
	)

	if err := s.db.QueryRow(ctx, selectQuery, id).Scan(
		&item.Id,
		&item.Type,
		&item.Name,
		&conditionId,
		&item.WorkplaceId,
		&item.UpdatedAt,
		&item.CreatedAt,
	); err != nil {
		s.logger.Warn(err.Error())
		return models.Item{}, err
	}
	condition := utills.IdsToConditions[conditionId]
	item.Condition = condition
	return item, nil
}

func (s *Storage) CreateItem(ctx context.Context, item models.Item) (int64, error) {
	query := `INSERT INTO resource_service.items (type, name, condition_id, workplace_id) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int64
	conditionId := utills.ConditionsToIds[item.Condition]
	if err := s.db.QueryRow(ctx, query, item.Type, item.Name, conditionId, item.WorkplaceId).Scan(&id); err != nil {
		s.logger.Warn(err.Error())
		return 0, err
	}

	return id, nil
}

func (s *Storage) UpdateItem(ctx context.Context, id int64, updateFields []Field) (models.Item, error) {
	updateQuery := `UPDATE resource_service.items SET `
	updateColumns, err := GenerateUpdates(itemsColumnsMap, updateFields)
	if err != nil {
		s.logger.Warn("cannot generate columns", err.Error())
		return models.Item{}, err
	}
	if len(updateColumns) == 0 {
		s.logger.Warn("not enough data to update")
		return models.Item{}, fmt.Errorf("not enough data to update")
	}
	updateQuery += updateColumns + fmt.Sprintf(" WHERE id = %d RETURNING *", id)
	var result models.Item
	if err := s.db.QueryRow(ctx, updateQuery).Scan(
		&result.Id,
		&result.Type,
		&result.Name,
		&result.Condition,
		&result.WorkplaceId,
		&result.UpdatedAt,
		&result.CreatedAt,
	); err != nil {
		s.logger.Warn("cannot update item", err.Error())
		return models.Item{}, err
	}

	return result, nil
}

func (s *Storage) DeleteItem(ctx context.Context, id int64) error {
	query := `DELETE FROM resource_service.items WHERE id = $1`
	if _, err := s.db.Exec(ctx, query, id); err != nil {
		s.logger.Warn("cannot delete item", err.Error())
		return err
	}
	return nil
}

func (s *Storage) GetItemsByWorkplace(ctx context.Context, workplaceId int64, page, pageSize int64) ([]models.Item, int64, error) {
	//TODO implement me
	panic("implement me")
}

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
