package models

var Category = &m{}

type m struct{}

type category struct {
	Id   int32  `json:"id"`
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func (*m) GetList(pid string) (list []category) {
	db.Where("pid = ?", pid).Find(&list)
	return
}

func (*m) GetTotal(pid string) (count int) {
	db.Model(&category{}).Where("pid = ?", pid).Count(&count)
	return
}
