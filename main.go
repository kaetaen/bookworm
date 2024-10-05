package main

import (
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/kaetaen/bookworm/config"
	ctrl "github.com/kaetaen/bookworm/controllers"
)

var Env map[string]string
var err error

func main() {
	config.LoadEnv()

	r := mux.NewRouter()

	user := ctrl.User{}

	r.HandleFunc("/user", user.Create).Methods("POST")

	http.ListenAndServe(":8000", r)
}
