package customtypes

// User ...
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// NotFoundError ...
type NotFoundError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Success ...
type Success struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
