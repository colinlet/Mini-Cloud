package controller

import "net/http"

var Api = api{}

type api struct{}

func (*api) Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("this is test~~"))
}
