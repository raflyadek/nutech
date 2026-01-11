package dto

type UserRegisterRequest struct {
	Id uint `json:"-"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	FirstName string `json:"first_name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	ProfileImage string `json:"profile_image"`
}

type UserLoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}