package entity

type User struct {
	Id      int32  `json:"id"`
	Openid  string `json:"openid"`
	Session string `json:"session"`
}
