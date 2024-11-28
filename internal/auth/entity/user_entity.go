package entity

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	FullName string
	Email    string
	Password string
}
