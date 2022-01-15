package gathererserver

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/imroc/req"
)

type kline struct {
	timestamp              string
	open, high, low, close float64
}

func getKline() {
	param := req.Param{
		"symbol":   configuration.Symbol,
		"interval": configuration.Interval,
	}
	r, err := req.Get(configuration.BinanceKlineUrl, param)

	if err != nil {
		log.Error(err)
	}

	klineString, err := r.ToString()
	if err != nil {
		log.Error(err)
		return
	}
	klineString = klineString[:len(klineString)-1]

	fmt.Println(klineString)
}
