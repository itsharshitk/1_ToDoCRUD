package model

type ErrorResponse struct {
	Status  int64  `json:"status"`
	Message string `json:"message"`
}
