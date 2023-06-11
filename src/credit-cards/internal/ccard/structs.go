package ccard

type CreditCard struct {
	id    int `json:"id"`
	Uid string `json:"uid"`
	CreditCardNumber string `json:"credit_card_number"`
	CreditCardExpiryDate string `json:"credit_card_expiry_date"`
	CreditCardType string `json:"credit_card_type"`
	
}
