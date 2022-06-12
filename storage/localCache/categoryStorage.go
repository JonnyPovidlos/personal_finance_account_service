package localCache

import (
	"personal_finance_account_service/models"
)

type userCategories struct {
	categories map[string]*models.Category
}

type categoryCacheStorage struct {
	counter         int
	usersCategories map[int]*userCategories
}

func newUserCategories() *userCategories {
	return &userCategories{
		categories: make(map[string]*models.Category),
	}
}

func NewCategoryCacheStorage() categoryCacheStorage {
	return categoryCacheStorage{0, make(map[int]*userCategories)}
}

func (c *categoryCacheStorage) Insert(category models.CreateCategory, userId int) (categoryId *int, err error) {
	userCategories, ok := c.usersCategories[userId]
	if !ok {
		userCategories = newUserCategories()
	}

	c.counter++
	categoryId = &c.counter
	userCategories.categories[category.Name] = &models.Category{
		Id:       *categoryId,
		Name:     category.Name,
		ParentId: category.ParentId,
		UserId:   userId,
	}
	c.usersCategories[userId] = userCategories
	//fmt.Println(*c.usersCategories[userId].categories["test"])
	return categoryId, nil
}
