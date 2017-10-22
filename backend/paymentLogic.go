package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"github.com/gorilla/mux"
	"time"
)

var ThePoolID string = "nil"

func makePayment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    var unprocessedPayment UnprocessedPayment
	_ = json.NewDecoder(r.Body).Decode(&unprocessedPayment)
	//fmt.Println(unprocessedPayment)

	var result Account
	
	if (params["collection"] == "loaners") {
		err := GetLoanersCollection().FindId(params["id"]).One(&result)
		errCheck(err)
		//fmt.Println(result)

		resp := createTempAccount(result.ID, unprocessedPayment)
		//fmt.Println(resp)
		if (resp != "") {
			sendPaymentToNessie(unprocessedPayment, resp)
			json.NewEncoder(w).Encode("Payment Processed")
		} else {
			http.Error(w, "Ivalid Account info!", http.StatusBadRequest)
		}

	} else if (params["collection"] == "donors") {
		err := GetDonorsCollection().FindId(params["id"]).One(&result)
		errCheck(err)
		
		resp := createTempAccount(result.ID, unprocessedPayment)
		if (resp != "") {
			sendPaymentToNessie(unprocessedPayment, resp)
			json.NewEncoder(w).Encode("Payment Processed")
		} else {
			http.Error(w, "Ivalid Account info!", http.StatusBadRequest)
		}

	} else {
		http.Error(w, "That collections does not exist!", http.StatusBadRequest)
		return
	}
}

func createPoolAccount() {
	account := CreateAccountObject("placeholder", "TheBig", "Pool", "bigmoney@youwishyouwhereme.com", "password", "1",  "no", "Merca", "FL", "69961")

 //-----------------------Create post request to Nessie-------------------------------------------------------

 	merchant := Merchant{"The Big Pool", []string{"Charity"}, account.Address, Geocode{1.0, 1.0}}
	var jsonStr, _ = json.Marshal(merchant)
	req, err := http.NewRequest("POST", "http://api.reimaginebanking.com/merchants?key=542922f7bba311ded255ef44e29df65f", bytes.NewBuffer(jsonStr))
	errCheck(err)
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	errCheck(err)
	defer resp.Body.Close()
	
	var response RspMerchant
	_ = json.NewDecoder(resp.Body).Decode(&response)
	fmt.Println(response)
	ThePoolID = response.ObjectCreated.ID
	account.ID = response.ObjectCreated.ID
  
 //-------------Save it to mongo DB----------------------------------------------------------------------------

	GetLoanersCollection().UpsertId(account.ID, account)
}

func sendPaymentToNessie(u UnprocessedPayment, id string) (string) {
	
	newPayment := Payment{ThePoolID, "balance", time.Now().String(), u.Amount, "Donation made to pool"}
	url := "http://api.reimaginebanking.com/accounts/" + id + "/purchases?key=542922f7bba311ded255ef44e29df65f"

	var jsonStr, _ = json.Marshal(newPayment)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	errCheck(err)
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	errCheck(err)
	defer resp.Body.Close()
	
	var webPage errorOb
	//fmt.Println(resp.Body)
	_ = json.NewDecoder(resp.Body).Decode(&webPage)
	fmt.Println(webPage)
	return "webPage"
}

func createTempAccount(id string, u UnprocessedPayment) (string) {	
	tempAccount := PaymentMeathod{"Credit Card", "temp", 0, 0, u.CreditCardNumber}

	url := "http://api.reimaginebanking.com/customers/" + id + "/accounts?key=542922f7bba311ded255ef44e29df65f"

	var jsonStr, _ = json.Marshal(tempAccount)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	errCheck(err)
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	//fmt.Println(resp.Body)
	errCheck(err)
	defer resp.Body.Close()
	
	var paymentMeathodRsp PaymentMeathodRsp
	_ = json.NewDecoder(resp.Body).Decode(&paymentMeathodRsp)
	//fmt.Println(paymentMeathodRsp)
	if (paymentMeathodRsp.Code < 400) { return paymentMeathodRsp.ObjectCreated.ID }
	return ""
}