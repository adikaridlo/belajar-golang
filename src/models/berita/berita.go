package artikel

import (
	"fmt"
	"net/http"
)

func Politik(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Artikel Politik")
}