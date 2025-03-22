package models

import "time"

type Item struct {
	Id          int64
	Type        string
	Name        string
	Condition   string
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
	Items             []Item
}

type TimeSlot struct {
	From   time.Time
	To     time.Time
	Reason string
}

type ParkingPlace struct {
	Id          int64
	Number      int64
	Location    string
	Type        string
	IsAvailable bool
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
