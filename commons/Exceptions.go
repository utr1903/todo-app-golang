package commons

// Exception : For returning meaningful responses in case of failure
type Exception struct {
	Code    string
	Message string
	Error   *error
}

// --- Not Found Exceptions ---

// UserNotFound : User is not available in database
func UserNotFound() *Exception {
	return &Exception{
		Code:    "UserNotFound",
		Message: "User not found!",
		Error:   nil,
	}
}

// ----------------------------

// --- Not Valid Exceptions ---

// RequestNotValid : Request structure is not in desired format
func RequestNotValid() *Exception {
	return &Exception{
		Code:    "RequestNotValid",
		Message: "Request structure not valid!",
		Error:   nil,
	}
}

// TokenNotValid : Token is not valid
func TokenNotValid() *Exception {
	return &Exception{
		Code:    "TokenNotValid",
		Message: "Token not valid!",
		Error:   nil,
	}
}

// UserIDNotValid : No User ID could be parsed from request
func UserIDNotValid() *Exception {
	return &Exception{
		Code:    "UserIDNotValid",
		Message: "User ID not valid!",
		Error:   nil,
	}
}

// ----------------------------
