package dtos

// TodoList : TodoList DB model
type TodoList struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Name   string `json:"name"`
}

// GetTodoLists : GetTodoLists result dto
type GetTodoLists struct {
	ListID   string `json:"listId"`
	ListName string `json:"listName"`
}
