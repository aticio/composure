package gathererserver

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/imroc/req"
)

type kline struct {
	timestamp              int
	open, high, low, close float64
}

func getKline() ([]kline, error) {
	param := req.Param{
		"symbol":   configuration.Symbol,
		"interval": configuration.Interval,
	}
	r, err := req.Get(configuration.BinanceKlineUrl, param)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	klines := []kline{}
	v := [][]interface{}{}
	r.ToJSON(&v)

	for _, k := range v {
		timestamp := int(k[0].(float64))
		open, _ := strconv.ParseFloat(k[1].(string), 64)
		high, _ := strconv.ParseFloat(k[2].(string), 64)
		low, _ := strconv.ParseFloat(k[3].(string), 64)
		close, _ := strconv.ParseFloat(k[4].(string), 64)
		k := kline{timestamp, open, high, low, close}
		klines = append(klines, k)
	}
	return klines, nil
}
