package models

type HotelItem struct {
	Address     string  `json:"address"`
	CountryCode string  `json:"country_code"`
	HotelCode   string  `json:"hotel_code"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	Name        string  `json:"name"`
}
