package dtos

// User : User DB model
type User struct {
	ID       string `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}
