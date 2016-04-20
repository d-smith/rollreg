package handlers
import (
	"net/http"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func AppsHandler(rw http.ResponseWriter, req *http.Request) {
	//Get apps for a user
	log.Info("apps query handler")
	rw.Write([]byte("apps query handler at your service\n"))
}

func PostAppsHandler(rw http.ResponseWriter, req *http.Request) {
	log.Info("post apps handler")
	rw.Write([]byte("post apps handler at your service\n"))
}

func GetApplicationByIdHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	log.Info("get app by id for ", vars["client_id"])
	rw.Write([]byte("get apps handler at your service\n"))
}

func PutApplicationByIdHandler(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	log.Info("put app by id for ", vars["client_id"])
	rw.Write([]byte("put apps handler at your service\n"))
}