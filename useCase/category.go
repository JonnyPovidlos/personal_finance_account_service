package useCase

import "personal_finance_account_service/models"

type CategoryStorage interface {
	Insert(category models.CreateCategory, userId int) (categoryId *int, err error)
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
