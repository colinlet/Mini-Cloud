package models

var Goods = &goodsModel{}

type goodsModel struct{}

type goods struct {
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

func (*goodsModel) GetList(pageNum int, pageSize int, maps interface{}) (list []goods) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&list)
	return
}

func (*goodsModel) GetTotal(maps interface{}) (count int) {
	db.Model(&goods{}).Where(maps).Count(&count)
	return
}

func (*goodsModel) Get(maps interface{}) (data goods) {
	db.Where(maps).Find(&data)
	return
}
