package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"personal_finance_account_service/models"
	"personal_finance_account_service/storage/localCache"
	"personal_finance_account_service/useCase"
)

var userStorage = localCache.NewUserCacheStorage()
var userUseCase = useCase.NewUserUseCases(&userStorage)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var createUser models.CreateUser
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(body, &createUser)
		if err != nil {
			log.Println(err)
		}
		log.Print(createUser)
		userId, err := userUseCase.Create(createUser)
		if err != nil {
			//log.Println(err)
			fmt.Fprint(w, "err: ", err)
		} else {
			fmt.Fprint(w, "userId: ", *userId)
		}
	}
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var signUser models.SignUser
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		if err := json.Unmarshal(body, &signUser); err != nil {
			log.Println(err)
		}
		log.Println(signUser)
		userId, err := userUseCase.Auth(signUser)
		if err != nil {
			log.Println(err)
		} else {
			w.Header().Set("userId", fmt.Sprintf("%d", *userId))
			fmt.Fprint(w, "userId: ", *userId)
		}

	}
}
