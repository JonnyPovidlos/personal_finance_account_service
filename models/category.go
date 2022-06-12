package models

import "fmt"

type Category struct {
	Id       int
	Name     string
	ParentId *int
	UserId   int
}

func (c *Category) String() string {
	return fmt.Sprintf(
		"Category: {Id: %d, Name: %s, ParentId: %d, UserId: %d}",
		c.Id, c.Name, *c.ParentId, c.UserId,
	)
}

type CreateCategory struct {
	Name     string
	ParentId *int
}

type EditCategory struct {
	Id       int
	Name     *string
	ParentId *int
}
