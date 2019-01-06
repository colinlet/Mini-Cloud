package models

import "Mini-Cloud/entity"

var Slideshow = &slideshowModel{}

type slideshowModel struct{}

func (*slideshowModel) table() string {
	return "mc_slideshow"
}

func (this *slideshowModel) GetList() (list []entity.Slideshow) {
	db.Table(this.table()).Where("status = ?", 1).Find(&list)
	return
}
