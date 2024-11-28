package model

type SuccessResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   any    `json:"data"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Error  any    `json:"error"`
}
