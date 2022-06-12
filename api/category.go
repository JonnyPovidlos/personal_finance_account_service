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
	"strconv"
	"strings"
)

var categoryStorage = localCache.NewCategoryCacheStorage()
var categoryUseCase = useCase.NewCategoryUseCase(&categoryStorage)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		userId, err := strconv.Atoi(r.Context().Value("UserId").(string))
		var createCategory models.CreateCategory
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(body, &createCategory)
		if err != nil {
			log.Println(err)
		}
		categoryId, err := categoryUseCase.Create(createCategory, userId)
		if err != nil {
			fmt.Fprint(w, "err: ", err)
		} else {
			fmt.Fprint(w, "categoryId: ", *categoryId)
		}
	}
}

func EditCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPatch {
		path := strings.Split(r.URL.String(), "/")
		categoryId, _ := strconv.Atoi(path[3])

		userId, err := strconv.Atoi(r.Context().Value("UserId").(string))
		var editCategory models.EditCategory
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(body, &editCategory)
		if err != nil {
			log.Println(err)
		}
		editCategory.Id = categoryId
		category := categoryUseCase.Edit(editCategory, userId)
		fmt.Fprint(w, "category: ", category.String())
	}
}
