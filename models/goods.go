package models

import "Mini-Cloud/entity"

var Goods = &goodsModel{}

type goodsModel struct{}

func (*goodsModel) GetList(pageNum int, pageSize int, maps interface{}) (list []entity.Goods) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&list)
	return
}

func (*goodsModel) GetTotal(maps interface{}) (count int) {
	db.Model(&entity.Goods{}).Where(maps).Count(&count)
	return
}

func (*goodsModel) Get(maps interface{}) (data entity.Goods) {
	db.Where(maps).Find(&data)
	return
}
