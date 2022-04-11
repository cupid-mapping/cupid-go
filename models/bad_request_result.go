package models

type BadRequestResult struct {
	// Error code for you to use programatically, for error logging and debugging.
	Code int32 `json:"code,omitempty"`
	// Error message intended to aid developers in debugging issues.
	Msg string `json:"msg,omitempty"`
}
