package main

type Account struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   struct {
		StreetNumber string `json:"street_number,omitempty"`
		StreetName   string `json:"street_name,omitempty"`
		City         string `json:"city,omitempty"`
		State        string `json:"state,omitempty"`
		Zip          string `json:"zip,omitempty"`
	} `json:"address,omitempty"`
}