package user

import "fmt"

type UserId int64

type User struct {
	Id UserId
	Name string
}

func NewUser(userId UserId) User {
	return User{
		Id:   userId,
		Name: fmt.Sprintf("%s_$v", "user_", userId),
	}
}
