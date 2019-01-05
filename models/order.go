package models

import "Mini-Cloud/entity"

var Order = &orderModel{}

type orderModel struct{}

func (*orderModel) GetList(pid string) (list []entity.Order) {
	db.Where("pid = ?", pid).Find(&list)
	return
}

func (*orderModel) GetTotal(pid string) (count int) {
	db.Model(&entity.Order{}).Where("pid = ?", pid).Count(&count)
	return
}
