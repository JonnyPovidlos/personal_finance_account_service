package localCache

import (
	"fmt"
	"personal_finance_account_service/models"
)

type UserCacheStorage struct {
	users map[int]*models.User
}

func NewUserCacheStorage() UserCacheStorage {
	return UserCacheStorage{make(map[int]*models.User)}
}

func (u *UserCacheStorage) Insert(createUser models.CreateUser) (id *int, err error) {
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
