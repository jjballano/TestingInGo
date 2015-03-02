package main 

import (	
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"learningGo/http/user"
	"net/http"	
	"io/ioutil"
	"strconv"
)

var users []interface{}
var added int = 0

func find(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id,_ := strconv.Atoi(vars["id"])
	json, err := json.MarshalIndent(users[id], "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(response, string(json))
}

func save(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	bodySave,_ := ioutil.ReadAll(request.Body)
	userReceived := user.New()
	json.Unmarshal(bodySave, &userReceived)
	users[added] = userReceived
	added++

	jsonValue := map[string]string{"Result":"OK"}
	result, _ := json.Marshal(jsonValue)
	fmt.Fprintf(response, string(result))
}

func notFound(response http.ResponseWriter, request *http.Request) {
	fmt.Println("ERROR")
	fmt.Fprintf(response, "")
}

func main() {
	users = make([]interface{},10)
	router := mux.NewRouter()
	router.HandleFunc("/find/{id}", find).Methods("GET")
	router.HandleFunc("/save", save).Methods("POST")
	router.HandleFunc("/{other}", notFound).Methods("GET", "POST")
	router.HandleFunc("/", notFound).Methods("GET", "POST")
	http.Handle("/", router)
	http.ListenAndServe("localhost:1337", nil)
}

func getJson() ([]byte, error){
	newUser := user.New()
	newUser.Username="Jesus"
	newUser.Email="jjballano@gmail.com"

	return json.MarshalIndent(newUser, "", "  ")
}