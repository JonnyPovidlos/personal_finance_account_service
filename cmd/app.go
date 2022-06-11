package main

import (
	"log"
	"net/http"
	"personal_finance_account_service/api"
)

var userPrefix = "/account"

func main() {
	http.HandleFunc(userPrefix+"/sign-up", api.SignUp)
	http.HandleFunc(userPrefix+"/sign-in", api.SignIn)

	log.Println("Server start")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
