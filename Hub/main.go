package main

import (
	"path/filepath"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/imroc/req"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	GathererAddress string
}

var configuration = Configuration{}

func init() {
	abs, err := filepath.Abs("./config.json")
	if err != nil {
		log.Error(err)
	}

	err = gonfig.GetConf(abs, &configuration)
	if err != nil {
		log.Error(err)
	}
}

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(get_data)
	s.StartAsync()
	s.StartBlocking()
}

func get_data() {
	r, err := req.Get(configuration.GathererAddress)
	if err != nil {
		log.Error(err)
	}
	log.Infof("%+v", r)
}
