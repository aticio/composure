package gathererserver

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/imroc/req"
)

type Price struct {
	Close []float64
}

func getPrice() (Price, error) {
	param := req.Param{
		"symbol":   configuration.Symbol,
		"interval": configuration.Interval,
		"limit":    configuration.Limit,
	}
	r, err := req.Get(configuration.BinanceKlineUrl, param)

	if err != nil {
		log.Error(err)
		return Price{}, err
	}

	p := Price{}
	v := [][]interface{}{}
	r.ToJSON(&v)

	for _, k := range v {
		c, _ := strconv.ParseFloat(k[4].(string), 64)
		p.Close = append(p.Close, c)
	}
	return p, nil
}
