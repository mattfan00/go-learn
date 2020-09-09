package main

import (
	"net/http"
	"webserver/helpers"
)

func main() {
	http.HandleFunc("/view/", helpers.MakeHandler(helpers.ViewHandler))
	http.HandleFunc("/edit/", helpers.MakeHandler(helpers.EditHandler))
	http.HandleFunc("/save/", helpers.MakeHandler(helpers.SaveHandler))
	http.ListenAndServe(":8080", nil)
}
