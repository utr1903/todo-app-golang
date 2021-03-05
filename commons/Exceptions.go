package commons

// Exception : For returning meaningful responses in case of failure
type Exception struct {
	M string
}

// --- Not Found Exceptions ---

// UserNotFound : User is not available in database
var UserNotFound string = "User not found!"

// ----------------------------

// --- Not Valid Exceptions ---

// RequestNotValid : Request structure is not in desired format
var RequestNotValid string = "Request structure not valid!"

// TokenNotValid : Token is not valid
var TokenNotValid string = "Token not valid!"

// UserIDNotValid : No User ID could be parsed from request
var UserIDNotValid string = "User ID not valid!"

// ----------------------------
