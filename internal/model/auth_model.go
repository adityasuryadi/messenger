package model

type RegisterRequest struct {
	Email                string `json:"email" validate:"required,unique=user"`
	Fullname             string `json:"fullname" validate:"required"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfitmation string `json:"password_confirmation" validate:"required,min=6,eqfield=Password"`
}