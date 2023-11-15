package dto

import "time"

type EjolRequest struct {
	Tid        string    `json:"tid" xml:"tid" form:"tid"`
	IpAddress  string    `json:"ip_address" xml:"ip_address" form:"ip_address"`
	DateInsert time.Time `json:"date_request" xml:"date_request" form:"date_request"`
}
