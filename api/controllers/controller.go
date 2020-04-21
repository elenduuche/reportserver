package controllers

import (
	"fmt"
	"log"
	"net/http"

	"dendrix.io/nayalabs/reportserver/api/internal/middlewares"
	"dendrix.io/nayalabs/reportserver/services"
	"github.com/gorilla/mux"
)

var (
	paymentController payment
)

const basePath = "/api/v1"

//StartUp sets up the controllers
func StartUp(data services.IDataService) http.Handler {
	rtr := mux.NewRouter()
	paymentController.registerRoutes(basePath, rtr)
	paymentController.registerServices(data)
	rtr.Use(middlewares.LoggingMw)
	return rtr
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	log.Println(err)
	fmt.Fprintf(w, err.Error())
}
