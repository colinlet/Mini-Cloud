package models

var Category = &m{}

type m struct{}

type category struct {
	Id   int32  `json:"id"`
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
	Img  string `json:"img"`
}

func (*m) GetList(pageNum int, pageSize int, maps interface{}) (list []category) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&list)
	return
}

func (*m) GetTotal(maps interface{}) (count int) {
	db.Model(&category{}).Where(maps).Count(&count)
	return
}
