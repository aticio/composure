package balanceserver

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetBalance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var b bytes.Buffer
	b.WriteString("soon...")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, b.String())
}
