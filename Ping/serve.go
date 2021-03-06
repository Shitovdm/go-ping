package Ping

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ping struct {
	Router *mux.Router
}

func Serve(port string) {
	sA := ping{}
	sA.Router = mux.NewRouter()
	sA.Router.HandleFunc("/ping", sA.Ping).Methods("GET")
	http.ListenAndServe(fmt.Sprintf(":%s", port), sA.Router)
}

func (sA *ping) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
