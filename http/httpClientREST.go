package main 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"learningGo/http/user"
)

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

	var userReceived = user.New()

	err = json.Unmarshal(body, &userReceived)
	if err != nil {
		panic(err)
	}

	fmt.Println("User received => ", userReceived)


}