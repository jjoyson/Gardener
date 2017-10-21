package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
    "github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()

	testAddAccounts()
	
	//router.HandleFunc("/account", getAccounts).Methods("GET")
    router.HandleFunc("/account/{collection}/{id}", getAccount).Methods("GET")
    router.HandleFunc("/account/{collection}", createAccount).Methods("POST")
	router.HandleFunc("/account/{collection}/{id}", deleteAccount).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var result interface{}
	if (params["collection"] == "loaners") {
		err := GetLoanersCollection().FindId(params["id"]).One(&result)
		errCheck(err)
		fmt.Println(result)
		json.NewEncoder(w).Encode(result)
	} else if (params["collection"] == "donors") {
		err := GetDonorsCollection().FindId(params["id"]).One(&result)
		errCheck(err)
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(&Account{})
	}
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    var account Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	//TODO set accoutn id to id given by capital 1 here
	if (params["collection"] == "loaners") {
		responce, err := GetLoanersCollection().UpsertId(account.ID, account)
		errCheck(err)
		json.NewEncoder(w).Encode(responce)
	} else if (params["collection"] == "donors") {
		responce, err := GetDonorsCollection().UpsertId(account.ID, account)
		errCheck(err)
		json.NewEncoder(w).Encode(responce)
	} else {
		json.NewEncoder(w).Encode(&Account{})
	}
    json.NewEncoder(w).Encode(account)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if (params["collection"] == "loaners") {
		GetLoanersCollection().RemoveId(params["id"])
	} else if (params["collection"] == "donors") {
		GetDonorsCollection().RemoveId(params["id"])
	} 
}

//func getAccounts(w http.ResponseWriter, r *http.Request) {
//	item := GetLoanersCollection().FindId(1)
//	json.NewEncoder(w).Encode(item)
//}

/*func getATMS() {
	res, err := http.Get("http://api.reimaginebanking.com/atms?key=542922f7bba311ded255ef44e29df65f")
	errCheck(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	errCheck(err)
	fmt.Printf("%s\n", string(body))
}*/

func testAddAccounts() {
	ac1 := CreateAccountObject("1", "Max", "rides", "12",  "it", "for", "", "pop") 
	ac2 := CreateAccountObject("2", "dog", "his", "11",  "is", "", "looking", "Zip string") 
	ac3 := CreateAccountObject("3","cool", "Pony", "10",  "hello", "me", "your", "") 

	c := GetLoanersCollection()

	_, err = c.UpsertId(ac1.ID, ac2)
	errCheck(err)
	_, err = c.UpsertId(ac2.ID, ac2)
	errCheck(err)
	c.UpsertId(ac3.ID, ac3)	

	var result interface{}
	err := c.FindId("1").One(&result)
	errCheck(err)
	fmt.Println(result)

	c = GetDonorsCollection()
	c.UpsertId(ac1.ID, ac1)	
	c.UpsertId(ac2.ID, ac2)
	c.UpsertId(ac3.ID, ac3)	

	err = c.FindId("2").One(&result)
	errCheck(err)
	fmt.Println(result)

}

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}