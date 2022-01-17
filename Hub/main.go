package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/imroc/req"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	GathererAddress string
	PearsonAddress  string
}

var configuration = Configuration{}

type price struct {
	Close []float64
}

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
	p, err := getData()
	if err != nil {
		return
	}

	pr, err := calculatePearsonsR(p)
	if err != nil {
		return
	}

	fmt.Println(pr)
}

func getData() (price, error) {
	r, err := req.Get(configuration.GathererAddress)
	if err != nil {
		log.Error(err)
		return price{}, err
	}

	p := price{}
	r.ToJSON(&p)
	return p, nil
}

func calculatePearsonsR(p price) (float64, error) {
	pb, err := json.Marshal(p)
	if err != nil {
		log.Error("Error creating post request to perason")
		return 0, err
	}
	r, err := req.Post(configuration.PearsonAddress, req.BodyJSON(pb))

	if err != nil {
		log.Error(err)
		return 0, err
	}

	prs, err := r.ToString()
	if err != nil {
		log.Error(err)
		return 0, err
	}

	pr, err := strconv.ParseFloat(prs, 64)
	if err != nil {
		log.Error(err)
		return 0, nil
	}

	return pr, nil
}
