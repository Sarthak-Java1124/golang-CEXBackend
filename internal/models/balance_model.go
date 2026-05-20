package models

type BalanceModel struct {
	ID string `json:"id"`

	UserId   string `json:"user_id"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Asset    string `json:"asset"`
}
