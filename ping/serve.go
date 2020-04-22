package ping

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Ping struct {
	Router *mux.Router
}

func Serve(port string) {
	sA := Ping{}
	sA.Router = mux.NewRouter()
	sA.Router.HandleFunc("/ping", sA.Ping).Methods("GET")
	http.ListenAndServe(fmt.Sprintf(":%s", port), sA.Router)
}

func (sA *Ping) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
