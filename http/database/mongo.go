package database

import (
	"gopkg.in/mgo.v2"
)

type Mongo struct{}

type Model interface{
	Save() bool
}

func (mongo *Mongo) Save(obj Model, collectionName string) bool {
	session, err := mgo.Dial("localhost")
	panicIfError(err)

	defer session.Close()

	database := session.DB("golangTests")
	collection := database.C(collectionName)
	err = collection.Insert(obj)
	panicIfError(err)
	return true
}

func panicIfError(err error){
	if err != nil{
		panic(err)
	}
}