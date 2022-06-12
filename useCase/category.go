package useCase

import "personal_finance_account_service/models"

type CategoryStorage interface {
	Insert(category models.CreateCategory, userId int) (categoryId *int, err error)
	Edit(category models.EditCategory, userId int) (*models.Category, error)
	Delete(categoryId int, userId int) error
}

type CategoryUseCase struct {
	storage CategoryStorage
}

func NewCategoryUseCase(storage CategoryStorage) *CategoryUseCase {
	return &CategoryUseCase{storage: storage}
}

func (c *CategoryUseCase) Create(createCategory models.CreateCategory, userId int) (*int, error) {
	categoryId, err := c.storage.Insert(createCategory, userId)
	if err != nil {
		return nil, err
	}
	return categoryId, nil
}

func (c *CategoryUseCase) Edit(editCategory models.EditCategory, userId int) (*models.Category, error) {
	category, err := c.storage.Edit(editCategory, userId)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryUseCase) Delete(categoryId int, userId int) error {
	return c.storage.Delete(categoryId, userId)
}
