package models

type Location struct {
	Id        uint     `json:"location_id"`
	IPAddress string   `json:"ip_address"`
	City      string   `json:"city"`
	Region    string   `json:"region"`
	Lat       string   `json:"lat"`
	Lon       string   `json:"lon"`
	Weather   *Weather `json:"weather" gorm:"foreignKey:Id"`
}
