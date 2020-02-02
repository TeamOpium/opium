package opium

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {


	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/ad", getAdList).Methods(http.MethodGet, http.MethodPost)
	api.HandleFunc("/ad/id/{adId:[0-9]+}", getAdById).Methods(http.MethodGet, http.MethodPost)
	api.HandleFunc("/ad/create", adCreate).Methods(http.MethodPost)

	log.Println("Starting")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}
