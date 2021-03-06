package commons

import "database/sql"

// CommonUtils : Central class for calling commonly used functions
type CommonUtils struct {
	Db *sql.DB
}

// DoesUserExist : Checks whether a user with given ID exists
func (c *CommonUtils) DoesUserExist(db *sql.DB, userID *string) bool {
	q := "select id from users where id = ?"
	var userExists string
	err := db.QueryRow(q, userID).Scan(&userExists)
	if err != nil {
		return false
	}

	return true
}

// DoesListExist : Checks whether a list with given ID belonging to the given user exists
func (c *CommonUtils) DoesListExist(db *sql.DB, listID *string, userID *string) bool {
	q := "select Id from lists where Id = ? and UserId = ?"
	var listExists string
	err := db.QueryRow(q, listID, userID).Scan(&listExists)
	if err != nil {
		return false
	}

	return true
}

// DoesItemExist : Checks whether an item with given ID belonging to the given user exists
func (c *CommonUtils) DoesItemExist(db *sql.DB, itemID *string, userID *string) bool {
	q := "select items.Id from items" +
		" join lists on items.ListId = lists.Id" +
		" where items.Id = ? and lists.UserId = ?"

	var itemExists string
	err := db.QueryRow(q, itemID, userID).Scan(&itemExists)
	if err != nil {
		return false
	}

	return true
}
