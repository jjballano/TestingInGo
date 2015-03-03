package database

import (
	//"gopkg.in/mgo.v2"
	"fmt"
)

type Mongo struct{}

type Model interface{
	Save() bool
}

func (mongo *Mongo) Save(obj Model, collectionName string) bool {
	fmt.Println("LLEGAAAA")
	return true
	/*
	session, err := mgo.Dial("localhost")
	panicIfError(err)

	defer.session.Close()

	database := session.DB("golangTests")
	collection := database.C(collectionName)



}

func panicIfError(err Error){
	if err != nil{
		panic(err)
	}*/	
}