package main

import (
	"net/http"

	server "composure/Dealer/domestic/dealerserver"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := httprouter.New()
	router.POST("/getDeal", server.GetDeal)
	log.Info("starting Dealer server on port 8084")
	log.Fatal(http.ListenAndServe(":8084", router))
}
