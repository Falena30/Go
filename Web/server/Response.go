package server

import (
	"fmt"
	"net/http"
)

//Index adalah fungsi untuk menampikan Apa akabar dengan menerima rspone dan request
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa Kaba!")
}
