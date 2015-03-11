package main 

import (	
	"fmt"
	"github.com/gorilla/mux"
	"net/http"	
	"os"
)

func sayHello(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintf(response, "Hello "+vars["name"])
}

func notFound(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Error")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", sayHello).Methods("GET")
	router.HandleFunc("/", notFound).Methods("GET", "POST")
	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}