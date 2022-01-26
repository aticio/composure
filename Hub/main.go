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
	SlopeAddress    string
	BalanceAddress  string
}

var configuration = Configuration{}

type Price struct {
	Close []float64
}

type AccountInformation struct {
	MakerCommission  int       `json:"makerCommission"`
	TakerCommission  int       `json:"takerCommission"`
	BuyerCommission  int       `json:"buyerCommission"`
	SellerCommission int       `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	UpdateTime       int       `json:"updateTime"`
	AccountType      string    `json:"accountType"`
	Balances         []Balance `json:"balances"`
}

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
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

	lrs, err := calculateSlope(p)
	if err != nil {
		return
	}
	fmt.Println(lrs)

	a, err := getBalance()
	if err != nil {
		return
	}
	fmt.Println(a)
}

func getData() (Price, error) {
	r, err := req.Get(configuration.GathererAddress)
	if err != nil {
		log.Error(err)
		return Price{}, err
	}

	p := Price{}
	r.ToJSON(&p)
	return p, nil
}

func calculatePearsonsR(p Price) (float64, error) {
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

func calculateSlope(p Price) (float64, error) {
	pb, err := json.Marshal(p)
	if err != nil {
		log.Error("Error creating post request to slope")
		return 0, err
	}
	r, err := req.Post(configuration.SlopeAddress, req.BodyJSON(pb))

	if err != nil {
		log.Error(err)
		return 0, err
	}

	lrss, err := r.ToString()
	if err != nil {
		log.Error(err)
		return 0, err
	}

	lrs, err := strconv.ParseFloat(lrss, 64)
	if err != nil {
		log.Error(err)
		return 0, nil
	}

	return lrs, nil
}

func getBalance() (AccountInformation, error) {
	r, err := req.Get(configuration.BalanceAddress)
	if err != nil {
		log.Error(err)
		return AccountInformation{}, err
	}

	a := AccountInformation{}
	r.ToJSON(&a)
	return a, nil
}
