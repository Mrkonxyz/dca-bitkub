package model

type DepositHistoryResponse struct {
	TxnID    string   `json:"txn_id"`
	Currency string   `json:"currency"`
	Amount   float64  `json:"amount"`
	Status   string   `json:"status"`
	Time     UnixTime `json:"time"`
}
