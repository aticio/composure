package main

import (
	"net/http"

	server "composure/Gatherer/domestic/gathererserver"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := httprouter.New()
	router.GET("/calculatepr", server.GetData)
	log.Info("starting Pearson server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
