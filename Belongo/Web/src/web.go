package web

import (
"net/http"
"fmt"
)

func init() {
	http.HandleFunc("/", helloweb)
}


func helloweb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Web placeholder")
}