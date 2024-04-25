package dto

type RegisterRequest struct {
	UserName             string `json:"user_name" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}

type RegisterResponse struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
