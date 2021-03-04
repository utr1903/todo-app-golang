package dtos

// TodoList : TodoList DB model
type TodoList struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Name   string `json:"name"`
}
