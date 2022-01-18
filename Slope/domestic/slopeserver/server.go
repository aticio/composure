package slopeserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CalculateSlope(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("soon...")
}
