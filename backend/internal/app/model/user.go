package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Balance  int    `json:"balance"`
	Password string `json:"password"`
}
