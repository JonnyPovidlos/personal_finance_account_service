package useCase

import (
	"golang.org/x/crypto/bcrypt"
	"personal_finance_account_service/models"
)

type UserStorage interface {
	Insert(createUser models.CreateUser) (id *int, err error)
}

type UserUseCases struct {
	storage UserStorage
}

func NewUserUseCases(storage UserStorage) *UserUseCases {
	return &UserUseCases{storage: storage}
}

func (u *UserUseCases) Create(user models.CreateUser) (*int, error) {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPass)
	id, err := u.storage.Insert(user)
	if err != nil {
		return nil, err
	}
	return id, nil
}
