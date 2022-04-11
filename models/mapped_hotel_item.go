package models

type MappedHotelItem struct {
	Address string `json:"address,omitempty"`
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	// The ID from your inventory
	Id        int32  `json:"id,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	// The ID from the partner catalog
	MappedHotelId string `json:"mapped_hotel_id,omitempty"`
	Name          string `json:"name,omitempty"`
}
