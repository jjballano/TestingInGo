package main 

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Email string
}

func server(response http.ResponseWriter, request *http.Request) {
	json, err := getJson()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(response, string(json))
}


func main() {
	http.HandleFunc("/", server)
	http.ListenAndServe("localhost:1337", nil)
}

func getJson() ([]byte, error){
	user := User{Name:"Jesus", Email:"jjballano@gmail.com"}

	return json.MarshalIndent(user, "", "  ")
}