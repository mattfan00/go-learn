package router

import (
	"fmt"
	"net/http"
)

func userIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "whtsa poppin")
}
