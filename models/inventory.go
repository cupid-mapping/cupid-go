package models

import (
	"time"
)

type Inventory struct {
	// If true, this is the inventory used as a reference when mapping
	Active    bool      `json:"active,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Inventory ID
	Id            int32 `json:"id,omitempty"`
	InvalidHotels int32 `json:"invalid_hotels,omitempty"`
	// Process status: * -1 - INACTIVE. * 0 - PENDING. * 1 - PROCESSING. * 2 - DONE. * 3 - FAILED.
	MappingProcessStatusId int32     `json:"mapping_process_status_id,omitempty"`
	Name                   string    `json:"name,omitempty"`
	TotalHotels            int32     `json:"total_hotels,omitempty"`
	UpdatedAt              time.Time `json:"updated_at,omitempty"`
	ValidHotels            int32     `json:"valid_hotels,omitempty"`
}
