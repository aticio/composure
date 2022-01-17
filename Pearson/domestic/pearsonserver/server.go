package pearsonserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type price struct {
	Close []float64
}

func CalculatePR(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pb, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error("Error reading body")
	}

	p := price{}
	err = json.Unmarshal(pb, &p)
	if err != nil {
		log.Error("Error parsing response")
	}

	fmt.Println(p)

}
