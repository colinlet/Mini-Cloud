package entity

type Goods struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"img"`
	Desc        string `json:"desc"`
	Price       int32  `json:"price"`
	OriginPrice int32  `json:"origin_price"`
	CategoryId  int32  `json:"category_id"`
	Showcase    string `json:"showcase"`
	List        string `json:"list"`
}
