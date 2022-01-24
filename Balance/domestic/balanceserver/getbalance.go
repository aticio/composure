package balanceserver

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/imroc/req"
)

func getBalance() {
	ts := strconv.FormatInt(time.Now().UTC().Unix()*(1000), 10)
	payload := fmt.Sprintf("&timestamp=%v", ts)
	mac := hmac.New(sha256.New, []byte(api_secret))
	_, err := mac.Write([]byte(payload))
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
	}
	fmt.Println(r)
}
