package main

import (
	"net/http"

	server "composure/Slope/domestic/slopeserver"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := httprouter.New()
	router.POST("/calculateslope", server.CalculateSlope)
	log.Info("starting Pearson server on port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
