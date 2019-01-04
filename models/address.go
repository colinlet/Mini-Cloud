package models

var Address = &addressModel{}

type addressModel struct{}

type address struct {
	Id   int32  `json:"id"`
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func (*addressModel) GetList(pid string) (list []address) {
	db.Where("pid = ?", pid).Find(&list)
	return
}

func (*addressModel) GetTotal(pid string) (count int) {
	db.Model(&address{}).Where("pid = ?", pid).Count(&count)
	return
}
