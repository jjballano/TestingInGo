package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	//Properties must be capitalized. They will be stored in lowercase
	Name string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost")

	if err != nil {
		panic (err)
	}
	defer session.Close()

	database := session.DB("golangTests")
	usersCollection := database.C("users")

	err = usersCollection.Insert(&Person{Name:"Jhon", Phone:"0894454545"},
						   &Person{Name:"Carla", Phone:"0895545454"})
	if err != nil {
		panic (err)
	}

	result := &Person{}

	//Properties in query must be in lowercase
	usersCollection.Find(bson.M{"name":"Carla"}).One(&result)

	fmt.Printf("Phone of Carla => %s", result.Phone)
	fmt.Println()
	
}