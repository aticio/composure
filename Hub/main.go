package main

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/imroc/req"
	log "github.com/sirupsen/logrus"
)

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(get_data)
	s.StartAsync()
	s.StartBlocking()
}

func get_data() {
	r, err := req.Get("http://localhost:8080/getData")
	if err != nil {
		log.Error(err)
	}
	log.Infof("%+v", r)
}
