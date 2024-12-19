package model

type MessageRequest struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

type MessageResponse struct {
	From    string `json:"from"`
	Message string `json:"message"`
}
