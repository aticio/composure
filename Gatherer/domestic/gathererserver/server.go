package gathererserver

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetData(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var b bytes.Buffer
	b.WriteString("soon...")
	fmt.Fprint(w, b.String())
}
