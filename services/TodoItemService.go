package services

import "database/sql"

// TodoItem : TodoList model
type TodoItem struct {
	ID      string `json:"id"`
	ListID  string `json:"listId"`
	Content string `json:"content"`
}

// TodoItemService : Implementation of TodoItemService
type TodoItemService struct{}

// GetItems : Returns all items
func (tls *TodoItemService) GetItems(db *sql.DB) ([]TodoItem, error) {
	rows, err := db.Query("select * from items")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []TodoItem{}

	for rows.Next() {
		var item TodoItem
		if rows.Scan(&item.ID, &item.ListID, &item.Content) != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
