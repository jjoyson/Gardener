package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
)

func getPayAccount(w http.ResponseWriter, r *http.Request) {
	//needs id string, collection string

}

func getPaymentHistory(w http.ResponseWriter, r *http.Request) {
	//needs id string, collection string

}

func makePayment(w http.ResponseWriter, r *http.Request) {

}

func createPoolAccount() {
	account := CreateAccountObject ("placeholder", "TheBig", "Pool", "bigmoney@youwishyouwhereme.com", "password", "1",  "no", "Merca", "FL", "69961")

 //-----------------------Create post request to Nessie-------------------------------------------------------

 	toBeSent := NessieAccount{account.FirstName, account.LastName, account.Address}
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
	account = CreateAccountObject(response.ObjectCreated.ID, account.FirstName, account.LastName, account.Email, account.Password, account.StreetNumber,  account.StreetName, account.City, account.State, account.Zip)
  
 //-------------Save it to mongo DB----------------------------------------------------------------------------

	GetLoanersCollection().UpsertId(account.ID, account)
}

func createPayAccount(id string, collection string) {
	
}