package salonhandlers

import (
	"log"
	"net/http"
	"shutter/models"
	"shutter/restapi/operations/salon"

	"github.com/go-openapi/runtime/middleware"
)

// NewHomeHandler handles a request for getting an entry
func SalonStopHandler() salon.SalonStopHandler {
	return &StopSalon{}
}

type StopSalon struct {
}

// Handle the get entry request
func (h *StopSalon) Handle(params salon.SalonStopParams, jwt interface{}) middleware.Responder {
	_, err := http.Get("http://192.168.1.22/scada-vis/objects/setvalue?data=%7B%22address%22%3A%222%2F2%2F30%22%2C%22datatype%22%3A1%2C%22value%22%3A%221%22%2C%22type%22%3A%22text%22%2C%22update%22%3Afalse%7D")
	if err != nil {
		log.Fatal(err)
	}
	return &salon.SalonOpenOK{Payload: &models.Success{Message: "Salon Stop"}}
}
