package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
    "github.com/gorilla/mux"
)

var CUST_URL = "http://api.reimaginebanking.com/customers?key=542922f7bba311ded255ef44e29df65f"

func getAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var result interface{}
	if (params["collection"] == "loaners") {
		err := GetLoanersCollection().FindId(params["id"]).One(&result)
		errCheck(err)
		json.NewEncoder(w).Encode(result)
	} else if (params["collection"] == "donors") {
		err := GetDonorsCollection().FindId(params["id"]).One(&result)
		errCheck(err)
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(&Account{})
	}
}

func getEmailAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var result interface{}
	if (params["collection"] == "loaners") {
		err := GetLoanersCollection().Find(bson.M{"email": noIDAccount.Email}).One(&result)
		errCheck(err)
		json.NewEncoder(w).Encode(result)
	} else if (params["collection"] == "donors") {
		err := GetDonorsCollection().Find(bson.M{"email": noIDAccount.Email}).One(&result)
		errCheck(err)
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(&Account{})
	}
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    var noIDAccount NoIDAccount
	_ = json.NewDecoder(r.Body).Decode(&noIDAccount)

	if(Find(bson.M{"email": noIDAccount.Email}).Count() != 0){
		http.Error(w, "Email already used!", http.StatusBadRequest)
		return;
	}


 //-----------------------Create post request to Nessie-------------------------------------------------------

 	toBeSent := NessieAccount{noIDAccount.FirstName, noIDAccount.LastName, noIDAccount.Address}
	var jsonStr, _ = json.Marshal(toBeSent)
	fmt.Println(toBeSent)
	req, err := http.NewRequest("POST", CUST_URL, bytes.NewBuffer(jsonStr))
	errCheck(err)
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	errCheck(err)
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	
	var response ResponseAccount
	_ = json.NewDecoder(resp.Body).Decode(&response)
	fmt.Println(response)
	account := CreateAccountObject(response.ObjectCreated.ID, noIDAccount.FirstName, noIDAccount.LastName, noIDAccount.Email, noIDAccount.Password, noIDAccount.StreetNumber,  noIDAccount.StreetName, noIDAccount.City, noIDAccount.State, noIDAccount.Zip)
  
 //-------------Save it to mongo DB----------------------------------------------------------------------------

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

/*func testAddAccounts() {
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

}*/