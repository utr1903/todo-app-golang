package commons

// Result : For returning meaningful responses in case of failure
type Result struct {
	Success bool        `json:"success"`
	Model   interface{} `json:"model"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Error   *error      `json:"error"`
}

// Success : General way to response with success
func Success(data interface{}, err error) *Result {
	return &Result{
		Success: true,
		Model:   data,
		Code:    "",
		Message: "",
		Error:   nil,
	}
}

// --- Not Found Exceptions ---

// UserNotFound : User is not available in database
func UserNotFound() *Result {
	return &Result{
		Success: false,
		Model:   nil,
		Code:    "UserNotFound",
		Message: "User not found!",
		Error:   nil,
	}
}

// ----------------------------

// --- Not Valid Exceptions ---

// RequestNotValid : Request structure is not in desired format
func RequestNotValid() *Result {
	return &Result{
		Success: false,
		Model:   nil,
		Code:    "RequestNotValid",
		Message: "Request structure not valid!",
		Error:   nil,
	}
}

// TokenNotValid : Token is not valid
func TokenNotValid() *Result {
	return &Result{
		Success: false,
		Model:   nil,
		Code:    "TokenNotValid",
		Message: "Token not valid!",
		Error:   nil,
	}
}

// UserIDNotValid : No User ID could be parsed from request
func UserIDNotValid() *Result {
	return &Result{
		Success: false,
		Model:   nil,
		Code:    "UserIDNotValid",
		Message: "User ID not valid!",
		Error:   nil,
	}
}

// ----------------------------
