package main

func CreateAccountObject (ID string, FirstName string, LastName string, StreetNumber string,  StreetName string, City string, State string, Zip string) (Account) {
	account := Account{
        ID: "",
		FirstName: "", 
		LastName: "",
        Address: Address{
			StreetNumber: "",
			StreetName: "",
			City: "",
			State: "",
			Zip: ""}}

	if (ID == "") {
		account.ID = "none"
	} else {
		account.ID = ID
	}

	if (FirstName  == "") {
		account.FirstName = "none"
	} else {
		account.FirstName = FirstName
	}

	if (LastName  == "") {
		account.LastName = "none"
	} else {
		account.LastName = LastName
	}

	if (StreetNumber == "") {
		account.StreetNumber = "none"
	} else {
		account.StreetNumber = StreetNumber
	}

	if (StreetName == "") {
		account.StreetName = "none"
	} else {
		account.StreetName = StreetName
	}

	if (City == "") {
		account.City = "none"
	} else {
		account.City = City
	}

	if (State == "") {
		account.State = "none"
	} else {
		account.State = State	
	}

	if (Zip == "") {
		account.Zip = "none"
	} else {
		account.Zip = Zip
	}
	
	return account;
}

type Account struct {
	ID string 
	FirstName string 
	LastName  string 
	Address
}

type Address   struct {
	StreetNumber string 
	StreetName   string 
	City         string 
	State        string 
	Zip          string 
} 