package dealerserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type BulkInfo struct {
	PriceData   Price
	PearsonsR   float64
	LinRegSlope float64
	AccountInfo AccountInformation
}

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

type Deal struct {
	Side string `json:"side"`
}

func GetDeal(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pbi, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error("Error reading body")
	}

	bi := BulkInfo{}
	err = json.Unmarshal(pbi, &bi)
	if err != nil {
		log.Error("Error parsing response")
	}

	getDeal(bi)

	var b bytes.Buffer
	b.WriteString("soon")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, b.String())
}
