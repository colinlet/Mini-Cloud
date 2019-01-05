package models

import "Mini-Cloud/entity"

var User = &userModel{}

type userModel struct{}

func (*userModel) GetByOpenid(openid string) (data entity.User) {
	db.Where("openid = ?", openid).Find(&data)
	return
}

func (*userModel) Insert(data *entity.User) {
	db.Create(&data)
}

func (*userModel) Update(id int32, maps map[string]interface{}) {
	db.Model(&entity.User{}).Where("id = ?", id).Update(maps)
}

func (*userModel) GetList(pid string) (list []entity.User) {
	db.Where("pid = ?", pid).Find(&list)
	return
}

func (*userModel) GetTotal(pid string) (count int) {
	db.Model(&entity.User{}).Where("pid = ?", pid).Count(&count)
	return
}
