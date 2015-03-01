package main 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name string
	Email string
}

func main() {
	response, err := http.Get("http://localhost:1337")
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var user User

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println("User received => ", user)


}