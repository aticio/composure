package balanceserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	BinanceAccountUrl string
	Base              string
	Quote             string
}

var configuration = Configuration{}
var api_key string
var api_secret string

func init() {
	abs, err := filepath.Abs("./config.json")
	if err != nil {
		log.Error(err)
	}

	err = gonfig.GetConf(abs, &configuration)
	if err != nil {
		log.Error(err)
	}

	api_key = os.Getenv("BINANCE_API_KEY")
	api_secret = os.Getenv("BINANCE_API_SECRET")
}

func GetBalance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	a, err := getBalance()
	if err != nil {
		var b bytes.Buffer
		b.WriteString("Error getting balance")
		fmt.Fprint(w, b.String())
	}

	pa, err := json.Marshal(a)
	if err != nil {
		var b bytes.Buffer
		b.WriteString("Error converting balance data to json")
		fmt.Fprint(w, b.String())
	}

	var b bytes.Buffer
	for _, p := range pa {
		b.WriteByte(p)
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, b.String())
}
