package models

type BalanceModel struct {
	ID string `json:"id"`

	UserId  string `json:"user_id"`
	Balance int    `json:"price"`

	AssetBalance map[string]int `json:"asset"`
}
