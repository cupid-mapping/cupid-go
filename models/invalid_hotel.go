package models

type InvalidHotel struct {
	Id     string   `json:"id,omitempty"`
	Reason []string `json:"reason,omitempty"`
}
