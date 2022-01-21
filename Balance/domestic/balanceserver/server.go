package balanceserver

import (
	"bytes"
	"fmt"
	"net/http"
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

func GetBalance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	getBalance()
	var b bytes.Buffer
	b.WriteString("soon...")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, b.String())
}
