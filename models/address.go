package models

import (
	"github.com/colinlet/Mini-Cloud/entity"
)

var Address = &addressModel{}

type addressModel struct{}

func (*addressModel) table() string {
	return "mc_address"
}

func (this *addressModel) Insert(data *entity.UserAddress) {
	db.Table(this.table()).Create(&data)
}

func (this *addressModel) UpdateStatus(mid int32, maps map[string]interface{}) {
	db.Table(this.table()).Where("mid = ?", mid).Update(maps)
}

func (this *addressModel) GetList(mid int32) (list []entity.UserAddress) {
	db.Table(this.table()).Where("mid = ?", mid).Find(&list)
	return
}

func (this *addressModel) Choose(id string) {
	db.Table(this.table()).Where("id = ?", id).Update(map[string]interface{}{"is_use": 1})
}

func (this *addressModel) GetCurrent(mid int32) (data entity.UserAddress) {
	db.Table(this.table()).Where("mid = ? AND is_use = ?", mid, 1).Find(&data)
	return
}
