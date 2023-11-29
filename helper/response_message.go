package helper

import "time"

const (
	Msg_StatusUnprocessableEntity = "invalid request, cannot be proceed."
	Msg_StatusInternalServerError = "internal server error"
	Msg_BadRequest                = "error validation / bad request"
	Msg_StatusNotFound            = "data not found."
	Msg_StatusOK                  = "OK"
)

func StringToDate(date string) time.Time {
	parsed, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return parsed
}
