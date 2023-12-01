package dto

type EjshipmentPingRequest struct {
	Tid       string `json:"tid" xml:"tid" form:"tid" validate:"omitempty,min=2,max=10,numeric"`
	IpAddress string `json:"ip_address" xml:"ip_address" form:"ip_address" validate:"omitempty,ip"`
}

type PingHostRequest struct {
	Url string
}

type PingHostResponse struct {
	Code    int
	Message interface{}
	Error   error
}
