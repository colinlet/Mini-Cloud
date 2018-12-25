package main

import (
	"Mini-Cloud/controller"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("web server start...")

	//web
	http.Handle("/wechat/", http.StripPrefix("/wechat/", http.FileServer(http.Dir("./images"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	//router
	http.HandleFunc("/api/", controller.Api.Run)

	go http.ListenAndServe(":80", nil)

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	}
}
