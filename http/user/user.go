package user

type User struct {
	Username string
	Email string
}

func New()(User) {
	return User{}
}