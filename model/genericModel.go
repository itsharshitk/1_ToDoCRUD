package model

type APIResponse struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`

	Data interface{} `json:"data,omitempty"`

	ErrorCode string      `json:"error_code,omitempty"`
	Details   interface{} `json:"error_details,omitempty"`
}
