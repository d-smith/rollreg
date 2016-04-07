package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/xtracdev/rollreg/handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/developers", handlers.DevsQueryHandler)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Warn(err.Error())
	}
}
