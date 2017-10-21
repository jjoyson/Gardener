package main

import (
	"labix.org/v2/mgo"
)

var MONGODB_URL string = "http://192.168.99.100:32768"
var DB_NAME string = "Gardener"
var LOANERS_NAME string = "Loaners"
var DONORS_NAME string = "Donors"

var mongoSession, _ = mgo.Dial(MONGODB_URL)

func GetLoanersCollection() (*mgo.Collection) {
	return mongoSession.DB(DB_NAME).C(LOANERS_NAME)
}

func GetDonorsCollection() (*mgo.Collection) {
	return mongoSession.DB(DB_NAME).C(DONORS_NAME)
}
