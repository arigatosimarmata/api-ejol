package dto

import "time"

type EjolDBRequest struct {
	IdxEjlog  int    `json:"idx_ejlog" xml:"idx_ejlog" form:"idx_ejlog" validate:"omitempty,numeric"`
	Tid       string `json:"tid,omitempty" xml:"tid" form:"tid" validate:"omitempty,min=2,max=10,numeric"`
	IpAddress string `json:"ip_address,omitempty" xml:"ip_address" form:"ip_address" validate:"omitempty,ip"`
	StartDate string `json:"start_date,omitempty" xml:"start" form:"start" validate:"omitempty,DateTimeWithDash"`
	EndDate   string `json:"end_date,omitempty" xml:"end" form:"end" validate:"omitempty,DateTimeWithDash"`
	Limit     int    `json:"limit,omitempty" xml:"limit" form:"limit" validate:"omitempty,numeric,lt=100000"`
	Page      int    `json:"page,omitempty" xml:"page" form:"page" validate:"numeric,gt=0"`
}

type EjolDBResponse struct {
	IdxEjlog  int       `json:"idx_ejlog"`
	Ejlog     string    `json:"ejlog"`
	CreatedAt time.Time `json:"created_at"`
}

type Pagination struct {
	StartPage int `json:"start_page"`
	LastPage  int `json:"last_page"`
	Total     int `json:"total"`
}

type CustomTime struct {
	time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02 15:04:05"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}
