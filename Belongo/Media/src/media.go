package media
import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/", hellomedia)
}


func hellomedia(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Media placeholder")
}