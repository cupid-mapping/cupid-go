package models

type MapHotelsResponse struct {
	Code          int32             `json:"code,omitempty"`
	Data          []MappedHotelItem `json:"data,omitempty"`
	InvalidHotels []InvalidHotel    `json:"invalid_hotels,omitempty"`
}
