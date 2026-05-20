package models

type User struct {
	ID string `json:"id"`
	Email string `json:"email"`
    Password string `json:"password"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
}
