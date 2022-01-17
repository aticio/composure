package main

import (
	"net/http"

	server "composure/Pearson/domestic/pearsonserver"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := httprouter.New()
	router.POST("/calculatepr", server.CalculatePR)
	log.Info("starting Pearson server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
