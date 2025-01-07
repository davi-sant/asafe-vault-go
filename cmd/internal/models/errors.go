package models

type ResponseError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Details any    `json:"details"`
}
