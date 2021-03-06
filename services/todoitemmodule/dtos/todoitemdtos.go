package dtos

// TodoItem : TodoList DB model
type TodoItem struct {
	ID      string `json:"id"`
	ListID  string `json:"listId"`
	Content string `json:"content"`
}

// GetTodoItems : GetTodoItems result dto
type GetTodoItems struct {
	ItemID  string `json:"itemId"`
	Content string `json:"content"`
}
