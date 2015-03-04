package user

import (
	"learningGo/http/database"
)

type User struct {
	Username string
	Email string
}

func New()(User) {
	return User{}
}

func (user *User) Save()(bool){
	mongo := database.Mongo{}
	return mongo.Save(user,"users")
}