package models

type Category struct {
	Id       int
	Name     string
	ParentId *int
	UserId   int
}

type CreateCategory struct {
	Name     string
	ParentId *int
}
