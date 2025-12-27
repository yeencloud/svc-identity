package domain

type CreateUserParams struct {
	Mail     string `validate:"required,email"`
	Username string `validate:"required,min=4,max=20,alphanum,startswithalpha"`
	Password string `validate:"required,min=8,containsany=0123456789,containsany=!@#$%&*()-_=+.:?"`
}
