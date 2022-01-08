package main

import (
	"net/http"

	server "composure/Gatherer/domestic/gathererserver"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := httprouter.New()
	router.GET("/getData", server.GetData)
	log.Info("starting Gatherer server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
