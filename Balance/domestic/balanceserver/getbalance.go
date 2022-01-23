package balanceserver

import (
	"fmt"
	"time"

	"github.com/imroc/req"
)

func getBalance() {
	param := req.Param{
		"timestamp":  time.Now().UnixMilli(),
		"recvWindow": 5000,
		"signature":  api_secret,
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
