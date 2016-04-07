package handlers

import "net/http"
import (
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func DevsQueryHandler(rw http.ResponseWriter, req *http.Request) {
	log.Info("devs query handler")
	rw.Write([]byte("devs query handler at your service\n"))
}

func PutDevHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	log.Info("put handler, email ", vars["email"])
	rw.Write([]byte("devs put handler at your service\n"))
}

func GetDevHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	log.Info("get handler, email ", vars["email"])
	rw.Write([]byte("devs get handler at your service\n"))
}
