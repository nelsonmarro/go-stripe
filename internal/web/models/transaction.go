package models

type Transaction struct {
	TransactionStatusID int    `json:"transaction_status_id"`
	Amount              int    `json:"amount"`
	Currency            string `json:"currency"`
	Latest4Digits       string `json:"latest_4_digits"`
	BankReturnCode      string `json:"bank_return_code"`
}
