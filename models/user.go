package models

import (
	"Mini-Cloud/entity"
)

var User = &userModel{}

type userModel struct{}

func (*userModel) table() string {
	return "mc_user"
}

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

func (this *userModel) GetInfo(session string) (data entity.UserInfo) {
	db.Table(this.table()).Where("session = ?", session).Scan(&data)
	return
}

func (this *userModel) GetBySession(session string) (data entity.User) {
	db.Where("session = ?", session).Find(&data)
	return
}

func (this *userModel) UpdateInfo(id int32, maps map[string]interface{}) {
	db.Table(this.table()).Where("id = ?", id).Update(maps)
}
