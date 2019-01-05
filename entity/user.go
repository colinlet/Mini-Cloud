package entity

type User struct {
	Id      int32  `json:"id"`
	Openid  string `json:"openid"`
	Session string `json:"session"`
}

type UserInfo struct {
	Id        int32  `json:"mid"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
}

type UserAddress struct {
	Id           int32  `json:"id"`
	Mid          int32  `json:"mid"`
	UserName     string `json:"user_name"`
	PostalCode   string `json:"postal_code"`
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	CountyName   string `json:"county_name"`
	DetailInfo   string `json:"detail_info"`
	NationalCode string `json:"national_code"`
	TelNumber    string `json:"tel_number"`
	IsUse        int32  `json:"is_use"`
}
