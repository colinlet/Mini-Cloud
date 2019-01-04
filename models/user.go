package models

var User = &userModel{}

type userModel struct{}

type user struct {
	Id   int32  `json:"id"`
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func (*userModel) GetList(pid string) (list []user) {
	db.Where("pid = ?", pid).Find(&list)
	return
}

func (*userModel) GetTotal(pid string) (count int) {
	db.Model(&user{}).Where("pid = ?", pid).Count(&count)
	return
}
