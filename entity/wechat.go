package entity

type Code2Session struct {
	Errcode    int32  `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}
