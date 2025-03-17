package database_models

import "time"

type Item struct {
	Id          int64
	Type        string
	Name        string
	ConditionId int64
	WorkplaceId int64
	Quantity    int64
	IsAvailable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Workplace struct {
	Id                int64
	Address           string
	Zone              string
	Floor             string
	Number            int64
	Type              string
	Capacity          int64
	Description       string
	IsAvailable       bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
	MaintenanceStatus string
	UniqueTag         string
}
