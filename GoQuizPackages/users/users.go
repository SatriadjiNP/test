package users

import "github.com/google/uuid"

type Users struct {
	userId   uuid.UUID
	userName string
}

func NewUsers(userName string) *Users {
	return &Users{
		userId:   uuid.New(),
		userName: userName,
	}
}
