package models

import "github.com/colinlet/Mini-Cloud/entity"

var Category = &categoryModel{}

type categoryModel struct{}

func (*categoryModel) GetList(pid string) (list []entity.Category) {
	db.Where("pid = ? AND status = ?", pid, 1).Find(&list)
	return
}

func (*categoryModel) GetTotal(pid string) (count int) {
	db.Model(&entity.Category{}).Where("pid = ?", pid).Count(&count)
	return
}
