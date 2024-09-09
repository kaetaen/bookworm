package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		teste := "Teste"

		fmt.Fprintf(w, "Hello, user! %s", teste)
	})

	http.ListenAndServe(":7000", r)

}
