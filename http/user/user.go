package user

type User struct {
	Name string
	Email string
}

func New()(User) {
	return User{}
}