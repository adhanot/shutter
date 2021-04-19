package sdbhandlers

import (
	"log"
	"net/http"
	"shutter/models"
	"shutter/restapi/operations/sdb"

	"github.com/go-openapi/runtime/middleware"
)

// NewHomeHandler handles a request for getting an entry
func SdbStopHandler() sdb.SdbStopHandler {
	return &StopSdb{}
}

type StopSdb struct {
}

// Handle the get entry request
func (h *StopSdb) Handle(params sdb.SdbStopParams, jwt interface{}) middleware.Responder {
	_, err := http.Get("http://192.168.1.22/scada-vis/objects/setvalue?data=%7B%22address%22%3A%222%2F2%2F60%22%2C%22datatype%22%3A1%2C%22value%22%3A%221%22%2C%22type%22%3A%22text%22%2C%22update%22%3Afalse%7D")
	if err != nil {
		log.Fatal(err)
	}
	return &sdb.SdbOpenOK{Payload: &models.Success{Message: "Sdb Stop"}}
}
