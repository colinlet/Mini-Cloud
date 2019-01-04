package models

var Order = &orderModel{}

type orderModel struct{}

type order struct {
	Id   int32  `json:"id"`
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func (*orderModel) GetList(pid string) (list []order) {
	db.Where("pid = ?", pid).Find(&list)
	return
}

func (*orderModel) GetTotal(pid string) (count int) {
	db.Model(&order{}).Where("pid = ?", pid).Count(&count)
	return
}
