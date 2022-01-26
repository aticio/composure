package slopeserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type Price struct {
	Close []float64
}

func CalculateSlope(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pb, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error("Error reading body")
	}

	p := Price{}
	err = json.Unmarshal(pb, &p)
	if err != nil {
		log.Error("Error parsing response")
	}

	slope := calculate(p)
	sSlope := fmt.Sprint(slope)

	var b bytes.Buffer
	b.WriteString(sSlope)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, b.String())
}
