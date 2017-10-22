package main

type Merchant struct {
	Name     string   `json:"name"`
	Category []string `json:"category"`
	Address `json:"address"`
	Geocode `json:"geocode"`
}

type Geocode struct {
	Lat int `json:"lat"`
	Lng int `json:"lng"`
} 

type RspMerchant struct {
	Message       string `json:"message"`
	Code          int    `json:"code"`
	ObjectCreated struct {
		Name     string   `json:"name"`
		Category []string `json:"category"`
		Address  struct {
			StreetNumber string `json:"street_number"`
			StreetName   string `json:"street_name"`
			City         string `json:"city"`
			State        string `json:"state"`
			Zip          string `json:"zip"`
		} `json:"address"`
		Geocode struct {
			Lat int `json:"lat"`
			Lng int `json:"lng"`
		} `json:"geocode"`
		CreationDate string `json:"creation_date"`
		ID           string `json:"_id"`
	} `json:"objectCreated"`
}

type Payment struct {
	MerchantID   string  `json:"merchant_id"`
	Medium       string  `json:"medium"`
	PurchaseDate string  `json:"purchase_date"`
	Amount       float64 `json:"amount"`
	Description  string  `json:"description"`
}

type PaymentMeathod struct {
	Type          string `json:"type"`
	Nickname      string `json:"nickname"`
	Rewards       int    `json:"rewards"`
	Balance       int    `json:"balance"`
	AccountNumber string `json:"account_number"`
}

type PaymentMeathodRsp struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	ObjectCreated struct {
		Type          string `json:"type"`
		Nickname      string `json:"nickname"`
		Rewards       int    `json:"rewards"`
		Balance       int    `json:"balance"`
		AccountNumber string `json:"account_number"`
		CustomerID    string `json:"customer_id"`
		ID            string `json:"_id"`
	} `json:"objectCreated,omitempty"`
}

type UnprocessedPayment struct {
	Amount float64 `json:"amount"`
	CreditCardNumber string `json:"credit_card_number"`
}

type errorOb struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Culprit []string `json:"culprit"`
}