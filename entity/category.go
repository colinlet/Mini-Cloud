package entity

type Category struct {
	Id   int32  `json:"id"`
	Pid  int32  `json:"pid"`
	Name string `json:"name"`
	Img  string `json:"img"`
}
