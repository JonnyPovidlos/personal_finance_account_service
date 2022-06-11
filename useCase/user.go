package useCase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"personal_finance_account_service/models"
)

type UserStorage interface {
	Insert(createUser models.CreateUser) (id *int, err error)
	GetByEmail(email string) (signUser *models.User, err error)
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

func (u *UserUseCases) Auth(signUser models.SignUser) (id *int, err error) {
	user, err := u.storage.GetByEmail(signUser.Email)
	if err != nil {
		return nil, err
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signUser.Password)); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, fmt.Errorf("incorrect pass or email")
		}
		return &user.Id, nil
	}
}
