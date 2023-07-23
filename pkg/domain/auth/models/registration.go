package models

type Registration struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username"`
	Password string `json:"password" validate:"required"`
}
