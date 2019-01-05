package models

import "Mini-Cloud/entity"

var User = &userModel{}

type userModel struct{}

func (*userModel) GetByOpenid(openid string) (data entity.User) {
	db.Where("openid = ?", openid).Find(&data)
	return
}

func (*userModel) Insert(maps interface{}) {
	db.NewRecord(maps)
}

func (*userModel) GetList(pid string) (list []entity.User) {
	db.Where("pid = ?", pid).Find(&list)
	return
}

func (*userModel) GetTotal(pid string) (count int) {
	db.Model(&entity.User{}).Where("pid = ?", pid).Count(&count)
	return
}
