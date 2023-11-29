package domain

import "time"

type CustomTime struct {
	time.Time
}

type Ejol struct {
	Ejlog         string `json:"ejlog"`
	LastFileEjlog string `json:"last_ejlog"`
}

func TableName() string {
	return ""
}
