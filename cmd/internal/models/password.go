package models

type Password struct {
	Id              int64  `json:"id"`
	UserId          int64  `json:"user_id"`
	ServiceName     string `json:"service_name"`
	ServiceUserName string `json:"service_user_name"`
	ServicePassword string `json:"service_password"`
}
