package localCache

import (
	"fmt"
	"personal_finance_account_service/models"
)

type userCacheStorage struct {
	users map[int]*models.User
}

func NewUserCacheStorage() userCacheStorage {
	return userCacheStorage{make(map[int]*models.User)}
}

func (u *userCacheStorage) Insert(createUser models.CreateUser) (id *int, err error) {
	for _, user := range u.users {
		if user.Email == createUser.Email {
			return nil, fmt.Errorf("already exists")
		}
	}
	id = new(int)
	*id = len(u.users) + 1
	u.users[*id] = &models.User{
		Id:        *id,
		FirstName: createUser.FirstName,
		LastName:  createUser.LastName,
		Email:     createUser.Email,
		Password:  createUser.Password,
	}
	return id, nil
}

func (u *userCacheStorage) GetByEmail(email string) (signUser *models.User, err error) {
	for _, user := range u.users {
		if user.Email == email {
			signUser = user
			break
		}
	}
	if signUser == nil {
		return nil, fmt.Errorf("incorrect pass or email")
	}
	return signUser, nil
}
