package models

type Result struct {
	Code int32        `json:"code,omitempty"`
	Data *interface{} `json:"data,omitempty"`
}

type InlineResponse200 struct {
	Result *Result     `json:"Result,omitempty"`
	Code   int32       `json:"code,omitempty"`
	Data   []Inventory `json:"data,omitempty"`
}

type InlineResponse2001 struct {
	Result *Result    `json:"Result,omitempty"`
	Code   int32      `json:"code,omitempty"`
	Data   *Inventory `json:"data,omitempty"`
}

type ConflictResult struct {
	// Error code for you to use programatically, for error logging and debugging.
	Code int32 `json:"code,omitempty"`
	// Error message intended to aid developers in debugging issues.
	Msg string `json:"msg,omitempty"`
}

type ServerErrorResult struct {
	// Error code for you to use programatically, for error logging and debugging.
	Code int32 `json:"code,omitempty"`
	// Error message intended to aid developers in debugging issues.
	Msg string `json:"msg,omitempty"`
}

type UnauthorizedErrorResult struct {
	// Error code for you to use programatically, for error logging and debugging.
	Code int32 `json:"code,omitempty"`
	// Error message intended to aid developers in debugging issues.
	Msg string `json:"msg,omitempty"`
}

type UploadInventoryResponse struct {
	InventoryId int32 `json:"catalog_id,omitempty"`
	Code        int32 `json:"code,omitempty"`
}

type NotFoundResult struct {
	// Error code for you to use programatically, for error logging and debugging.
	Code int32 `json:"code,omitempty"`
	// Error message intended to aid developers in debugging issues.
	Msg string `json:"msg,omitempty"`
}
