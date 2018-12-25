package controller

import (
	"io/ioutil"
	"net/http"
)

var Api = api{}

type api struct{}

func (*api) Run(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	act := ""
	if _, ok := params["act"]; ok {
		act = params["act"][0]
	}
	if len(act) < 1 {
		w.Write([]byte("error"))
		return
	}
	data, err := ioutil.ReadFile("./data/" + act + ".json")
	if err != nil {
		w.Write([]byte(act + " no exist"))
		return
	}
	w.Write(data)
}
