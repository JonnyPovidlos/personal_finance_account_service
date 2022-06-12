package localCache

import (
	"fmt"
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

func (c *categoryCacheStorage) Edit(category models.EditCategory, userId int) (models.Category, error) {
	userCategories, ok := c.usersCategories[userId]
	if !ok {
		userCategories = newUserCategories()
	}
	var name string
	for _, ctgr := range userCategories.categories {
		if category.Id == ctgr.Id {
			name = ctgr.Name
			if category.Name != nil {
				name = *category.Name
				delete(userCategories.categories, ctgr.Name)
				userCategories.categories[name] = &models.Category{
					Id:       ctgr.Id,
					Name:     *category.Name,
					ParentId: ctgr.ParentId,
					UserId:   userId,
				}
			}
			if category.ParentId != nil {
				userCategories.categories[name].ParentId = category.ParentId
			}
		}
	}
	return *userCategories.categories[name], nil
}

func (c *categoryCacheStorage) Delete(categoryId int, userId int) error {
	deleteCategory := c.getById(categoryId, userId)
	if deleteCategory != nil {
		delete(c.usersCategories[userId].categories, deleteCategory.Name)
		return nil
	}
	return fmt.Errorf("not exists category")
}

func (c *categoryCacheStorage) getById(categoryId int, userId int) *models.Category {
	userCategories, ok := c.usersCategories[userId]
	if !ok {
		userCategories = newUserCategories()
	}
	for _, category := range userCategories.categories {
		if category.Id == categoryId {
			return category
		}
	}
	return nil
}
