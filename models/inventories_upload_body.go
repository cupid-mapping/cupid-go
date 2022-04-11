package models

import (
	"os"
)

type InventoriesUploadBody struct {
	File              **os.File `json:"file"`                       // CSV File to be uploaded
	Name              string    `json:"name"`                       // Inventory name
	HeaderId          int32     `json:"header_id"`                  // Column index for hotel ID
	HeaderName        int32     `json:"header_name"`                // Column index for hotel name
	HeaderAddress     int32     `json:"header_address"`             // Column index for hotel address
	HeaderCity        int32     `json:"header_city,omitempty"`      // Column index for hotel city
	HeaderCountryCode int32     `json:"header_country_code"`        // Column index for hotel country code
	HeaderLatitude    int32     `json:"header_latitude,omitempty"`  // Column index for hotel latitude
	HeaderLongitude   int32     `json:"header_longitude,omitempty"` // Column index for hotel longitude
}
