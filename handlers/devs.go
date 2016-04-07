package handlers

import "net/http"
import log "github.com/Sirupsen/logrus"

func DevsQueryHandler(rw http.ResponseWriter, req *http.Request) {
	log.Info("devs query handler")
	rw.Write([]byte("devs query handler at your service\n"))
}
