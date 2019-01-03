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
}

func (*goodsModel) GetList(pageNum int, pageSize int, maps interface{}) (list []goods) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&list)
	return
}

func (*goodsModel) GetTotal(maps interface{}) (count int) {
	db.Model(&goods{}).Where(maps).Count(&count)
	return
}
