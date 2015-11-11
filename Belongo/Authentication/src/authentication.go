package authentication
import (
	"net/http"
	"fmt"
)


func init() {
	http.HandleFunc("/", helloAuth)
}


func helloAuth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Auth placeholder")
}