package main

import (
	"net/http"

	server "composure/Balance/domestic/balanceserver"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := httprouter.New()
	router.GET("/getBalance", server.GetBalance)
	log.Info("starting Balance server on port 8083")
	log.Fatal(http.ListenAndServe(":8083", router))
}
