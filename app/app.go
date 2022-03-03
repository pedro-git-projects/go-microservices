package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port string = ":8080"

func Start() {

	router := mux.NewRouter()

	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	log.Print("starting application on port" + port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request recieved")
}
