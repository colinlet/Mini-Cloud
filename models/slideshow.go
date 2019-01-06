package models

import "Mini-Cloud/entity"

var Slideshow = &slideshowModel{}

type slideshowModel struct{}

func (*slideshowModel) table() string {
	return "mc_slideshow"
}

func (this *slideshowModel) GetList() (list []entity.Slideshow) {
	db.Table(this.table()).Find(&list)
	return
}
