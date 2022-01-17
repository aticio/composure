package gathererserver

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/imroc/req"
)

type price struct {
	Close []float64
}

func getPrice() (price, error) {
	param := req.Param{
		"symbol":   configuration.Symbol,
		"interval": configuration.Interval,
	}
	r, err := req.Get(configuration.BinanceKlineUrl, param)

	if err != nil {
		log.Error(err)
		return price{}, err
	}

	p := price{}
	v := [][]interface{}{}
	r.ToJSON(&v)

	for _, k := range v {
		c, _ := strconv.ParseFloat(k[4].(string), 64)
		p.Close = append(p.Close, c)
	}
	return p, nil
}
