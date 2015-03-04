package main 

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"learningGo/http/user"
	"bytes"
)

func main() {
	
	createUser()


	//findUser()

}

func createUser() {
	buffer := bytes.NewReader([]byte(`{"username":"jjballano", "email": "jjballano@gmail.com"}`))
	responseSave, err := http.Post("http://localhost:1337/save","application/json",buffer)
	if err != nil {
		panic(err)
	}
	defer responseSave.Body.Close()
	bodySave,_ := ioutil.ReadAll(responseSave.Body)
	fmt.Println("Response => ",string(bodySave))
}

func findUser() {
	response, err := http.Get("http://localhost:1337/find/0")
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