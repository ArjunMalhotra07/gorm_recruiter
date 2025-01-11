package models

type Response struct {
	Message string      `json:"message"`
	Jwt     *string     `json:"jwt"`
	Data    interface{} `json:"data"`
}
