package models

type Category struct {
	Id       int
	Name     string
	ParentId int
	UserId   int
}
