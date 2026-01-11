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

type UserUpdateProfileRequest struct {
	Id uint `json:"-"`
	Email string `json:"-"`
	Password string `json:"-"`
	FirstName string `json:"first_name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	ProfileImage string `json:"-"`
}

type UserUpdateImageRequest struct {
	Id uint `json:"-"`
	Email string `json:"-"`
	Password string `json:"-"`
	FirstName string `json:"-"`
	LastName string `json:"-"`
	ProfileImage string `json:"profile_image" validate:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserProfileResponse struct {
	Id uint `json:"-"`
	Email string `json:"email"`
	Password string `json:"-"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	ProfileImage string `json:"profile_image"`	
}