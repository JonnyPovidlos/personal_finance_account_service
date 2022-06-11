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
