package models

type Test struct {
	Model
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func GetTest(pageNum int, pageSize int, maps interface{}) (test []Test) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&test)
	return
}
