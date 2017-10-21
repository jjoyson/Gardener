package main

import (
	//"fmt"
	//"io/ioutil"
	"net/http"
	//"encoding/json"
    "log"
    "github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/account", getAccounts).Methods("GET")
    router.HandleFunc("/account/{id}", getAccount).Methods("GET")
    router.HandleFunc("/account/{id}", createAccount).Methods("POST")
	router.HandleFunc("/account/{id}", deleteAccount).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

/*func getATMS() {
	res, err := http.Get("http://api.reimaginebanking.com/atms?key=542922f7bba311ded255ef44e29df65f")
	errCheck(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	errCheck(err)
	fmt.Printf("%s\n", string(body))
}*/

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