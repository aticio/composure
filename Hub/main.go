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
	Gatherer_Address string
}

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(get_data)
	s.StartAsync()
	s.StartBlocking()
}

func get_data() {
	abs, err := filepath.Abs("./config.json")
	if err != nil {
		log.Error(err)
	}

	configuration := Configuration{}
	err = gonfig.GetConf(abs, &configuration)

	if err != nil {
		log.Error(err)
		panic(err)
	}

	r, err := req.Get(configuration.Gatherer_Address)
	if err != nil {
		log.Error(err)
	}
	log.Infof("%+v", r)
}
