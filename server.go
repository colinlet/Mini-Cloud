package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("web server start...")

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./images"))))
	go http.ListenAndServe(":8090", nil)

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	}
}
