package app

import (
	"log"
	"net/http"
)

var port string = ":8080"

func Start() {

	mux := http.NewServeMux()

	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/costumers", getAllCustomers)

	log.Print("starting application on port" + port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}

}
