package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//testAddAccounts()
	createPoolAccount() //This will create a account that will be the pool
	
	router.HandleFunc("/account/{collection}/{id}", getAccount).Methods("GET") //in accountLogic.go
	router.HandleFunc("/account/{collection}/{email}", getEmailAccount).Methods("GET") //in accountLogic.go
    router.HandleFunc("/account/{collection}", createAccount).Methods("POST") //in accountLogic.go
	router.HandleFunc("/account/{collection}/{id}", deleteAccount).Methods("DELETE") //in accountLogic.go
	
	//router.HandleFunc("/payments/{id}", getPaymentHistory).Methods("GET")
    router.HandleFunc("/payments/{collection}/{id}", makePayment).Methods("POST") //creditcards
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}