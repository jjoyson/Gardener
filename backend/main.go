package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	//"encoding/json"
    //"log"
    //"github.com/gorilla/mux"
)


func main() {
	//router := mux.NewRouter()

	testAddAccounts()
	
	//router.HandleFunc("/account", getAccounts).Methods("GET")
    //router.HandleFunc("/account/{id}", getAccount).Methods("GET")
    //router.HandleFunc("/account/{id}", createAccount).Methods("POST")
	//router.HandleFunc("/account/{id}", deleteAccount).Methods("DELETE")
	
	//log.Fatal(http.ListenAndServe(":8000", router))
}

/*func getATMS() {
	res, err := http.Get("http://api.reimaginebanking.com/atms?key=542922f7bba311ded255ef44e29df65f")
	errCheck(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	errCheck(err)
	fmt.Printf("%s\n", string(body))
}*/

func testAddAccounts() {
	ac := CreateAccountObject("1", "Max", "rides", "12",  "it", "for", "", "pop") 
	//ac2 := CreateAccountObject("2", "dog", "his", "11",  "is", "", "looking", "Zip string") 
	//ac3 := CreateAccountObject("3","cool", "Pony", "10",  "hello", "me", "your", "") 

	c := GetLoanersCollection()

	c.Insert(ac)
	/*c.UpsertId(ac2.ID, ac2)
	c.UpsertId(ac3.ID, ac3)*/	

	q := c.FindId(1)
	fmt.Println(q)

	/*c = GetDonorsCollection()
	c.Upsert(ac1.ID, ac1)	
	c.Upsert(ac2.ID, ac2)
	c.Upsert(ac3.ID, ac3)	

	q = c.FindId(2)
	fmt.Println(q)*/

}

func errCheck(e error) {
     if e != nil {
         panic(e)
     }
}

func getAccounts(w http.ResponseWriter, r *http.Request) {

}

func getAccount(w http.ResponseWriter, r *http.Request) {

}

func createAccount(w http.ResponseWriter, r *http.Request) {

}

func deleteAccount(w http.ResponseWriter, r *http.Request) {

}