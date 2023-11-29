package dto

type EjolNFSRequest struct {
	Tid         string `json:"tid" xml:"tid" form:"tid" validate:"omitempty,min=2,max=10,numeric"`
	IpAddress   string `json:"ip_address" xml:"ip_address" form:"ip_address" validate:"omitempty,ip"`
	DateRequest string `json:"date_request" xml:"date_request" form:"date_request" validate:"required,DateOnlyWithDash,DateLessThan"`
}
