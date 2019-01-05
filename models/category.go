package models

import "Mini-Cloud/entity"

var Category = &categoryModel{}

type categoryModel struct{}

func (*categoryModel) GetList(pid string) (list []entity.Category) {
	db.Where("pid = ?", pid).Find(&list)
	return
}

func (*categoryModel) GetTotal(pid string) (count int) {
	db.Model(&entity.Category{}).Where("pid = ?", pid).Count(&count)
	return
}
