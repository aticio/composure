package main

import (
	"fmt"
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
	s.Every(5).Seconds().Do(initOps)
	s.StartAsync()
	s.StartBlocking()
}

func initOps() {
	close, err := getData()
	if err != nil {
		return
	}

	fmt.Println(close)
}

func getData() ([]float64, error) {
	r, err := req.Get(configuration.GathererAddress)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var close []float64
	r.ToJSON(&close)
	return close, nil
}
