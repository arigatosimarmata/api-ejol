package dto

type EjshipmentSearchRequest struct {
	Kategori  string `json:"kategori" xml:"kategori" form:"kategori" validate:"omitempty,max=10"`
	Tid       string `json:"tid" xml:"tid" form:"tid" validate:"omitempty,min=2,max=10,numeric"`
	IpAddress string `json:"ip_address" xml:"ip_address" form:"ip_address" validate:"omitempty,ip"`
}

type SearchHostRequest struct {
	Url     string
	Request EjshipmentSearchRequest
}
type SearchHostResponse PingHostResponse
