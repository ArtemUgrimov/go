package users

import (
	"hw15/internal/routes"
	"math/rand"
)

type User struct {
	Id       int64
	Nickname string
	Routes   []routes.Route
}

func FromRequest(ur UserCreationRequest) User {
	user := User{
		Id:       rand.Int63(),
		Nickname: ur.Nickname,
	}
	user.Routes = make([]routes.Route, 0)
	return user
}
