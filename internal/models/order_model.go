package models

type OrderModel struct {
	ID                string `json:"id"`
	Market            string `json:"market"`
	Side              string `json:"side"`
	Type              string `json:"type"`
	UserId            string `json:"user_id"`
	Status            string `json:"status"`
	Price             int    `json:"price"`
	Quantity          int    `json:"quantity"`
	RemainingQuantity int    `json:"remaining_quantity"`
	CreatedAt         string `json:"created_at"`
}
