package gathererserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	BinanceKlineUrl string
	Symbol          string
	Interval        string
}

var configuration = Configuration{}

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

func GetData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	price, err := getPrice()
	if err != nil {
		var b bytes.Buffer
		b.WriteString("Error getting kline data")
		fmt.Fprint(w, b.String())
	}

	fmt.Println(price)

	pb, err := json.Marshal(price)
	if err != nil {
		var b bytes.Buffer
		b.WriteString("Error converting price object to json")
		fmt.Fprint(w, b.String())
	}

	var b bytes.Buffer
	for _, p := range pb {
		b.WriteByte(p)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, b.String())
}
