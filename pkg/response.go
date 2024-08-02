package pkg

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token      string `json:"token"`
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

type SignupRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

type User struct {
	ID       int
	Email    string
	Password string
}