package gathererserver

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/imroc/req"
)

func getKline() {
	param := req.Param{
		"symbol":   configuration.Symbol,
		"interval": configuration.Interval,
	}
	r, err := req.Get(configuration.BinanceKlineUrl, param)

	if err != nil {
		log.Error(err)
	}

	klines := interface{}(nil)
	r.ToJSON(&klines)
	fmt.Println(klines)
}
