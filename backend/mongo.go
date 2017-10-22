package main

import (
	"labix.org/v2/mgo"
)

var MONGODB_URL string = "mongo:27017"
var DB_NAME string = "Gardener"
var LOANERS_NAME string = "Loaners"
var DONORS_NAME string = "Donors"

var mongoSession, err = mgo.Dial(MONGODB_URL)

func GetLoanersCollection() (*mgo.Collection) {
	errCheck(err)
	return mongoSession.DB(DB_NAME).C(LOANERS_NAME)
}

func GetDonorsCollection() (*mgo.Collection) {
	errCheck(err)
	return  mongoSession.DB(DB_NAME).C(DONORS_NAME)
}
