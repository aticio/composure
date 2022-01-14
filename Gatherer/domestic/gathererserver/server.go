package gathererserver

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
	getKline()
	var b bytes.Buffer
	b.WriteString("soon...")
	fmt.Fprint(w, b.String())
}
