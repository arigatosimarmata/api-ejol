package domain

import "time"

type EjolDB struct {
	IdxEjlog  int    `json:"idx_ejlog"`
	Ejlog     string `json:"ejlog"`
	CreatedAt string `json:"created_at"`
}

type EjolDBRequest struct {
	IdxEjol   int
	Tid       string
	IpAddress string
	StartDate time.Time
	EndDate   time.Time
	Limit     int
	Kanwil    string
	Page      int
	TableName string
	DbName    string
}

type EjolDBResponse struct {
	Data        []EjolDB
	CurrentPage int
	LastPage    int
	Total       int
}
