package main

import (
	"log"
	"net/http"
	"personal_finance_account_service/api"
)

var userPrefix = "/api/account"
var categoryPrefix = "/api/category"

func main() {
	mux := http.NewServeMux()
	createCategoryHandler := http.HandlerFunc(api.CreateCategory)
	mux.Handle(categoryPrefix, api.CheckAuth(createCategoryHandler))
	mux.Handle(categoryPrefix+"/", api.CheckAuth(http.HandlerFunc(api.EditCategory)))

	mux.HandleFunc(userPrefix+"/sign-up", api.SignUp)
	mux.HandleFunc(userPrefix+"/sign-in", api.SignIn)

	log.Println("Server start")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
