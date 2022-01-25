package balanceserver

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/imroc/req"
	log "github.com/sirupsen/logrus"
)

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

func getBalance() (AccountInformation, error) {
	ts := strconv.FormatInt(time.Now().UTC().Unix()*(1000), 10)
	payload := fmt.Sprintf("&timestamp=%v", ts)
	mac := hmac.New(sha256.New, []byte(api_secret))
	_, err := mac.Write([]byte(payload))
	if err != nil {
		log.Error(err)
		return AccountInformation{}, err
	}

	param := req.QueryParam{
		"signature": hex.EncodeToString(mac.Sum(nil)),
		"timestamp": ts,
	}

	header := req.Header{
		"X-MBX-APIKEY": api_key,
	}
	r, err := req.Get(configuration.BinanceAccountUrl, header, param)

	if err != nil {
		log.Error(err)
		return AccountInformation{}, err
	}

	a := AccountInformation{}
	err = r.ToJSON(&a)

	if err != nil {
		log.Error(err)
		return AccountInformation{}, err
	}

	return a, nil
}
