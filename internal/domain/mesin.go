package domain

type Mesin struct {
	Tid       string
	IpAddress string
	Kanwil    string
}

func TableAtmMapping() string {
	return "atm_mappings"
}
