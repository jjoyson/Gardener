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
	ID string `json:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address  `json:"address"`
}

type Address   struct {
	StreetNumber string `json:"street_number"`
	StreetName   string `json:"street_name"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"` 
}

type NoIDAccount struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address  `json:"address"`
}

type ResponseAccount struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	ObjectCreated struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address    `json:"address"`
		ID string `json:"_id"`
	} `json:"objectCreated"`
}

