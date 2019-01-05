package entity

type User struct {
	Id      int32  `json:"id"`
	Openid  string `json:"openid"`
	Session string `json:"session"`
}

type UserInfo struct {
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
}
