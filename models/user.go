package models

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CreateUser struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type SignUser struct {
	Email    string
	Password string
}
