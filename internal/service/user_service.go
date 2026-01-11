package service

import (
	"fmt"
	"nutech-test/internal/dto"
	"nutech-test/internal/entity"
	"nutech-test/util"

	"golang.org/x/crypto/bcrypt"
)


type UserRepository interface {
	Create(user entity.User) error
	GetByEmail(email string) (entity.User, error)
	UpdateUserByEmail(user *entity.User) (entity.User, error)
	UpdateImageByEmail(user *entity.User) (entity.User, error)
}

type UserServ struct {
	userRepository UserRepository	
}

func NewUserService(ur UserRepository) *UserServ {
	return &UserServ{userRepository: ur}
}

func(us *UserServ) Create(req dto.UserRegisterRequest) error {
	//hash the request password
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hashing password %s", err)
	}

	user := entity.User{
		Email: req.Email,
		Password: string(passHash),
		FirstName: req.FirstName,
		LastName: req.LastName,
	}

	if err := us.userRepository.Create(user); err != nil {
		return fmt.Errorf("create user %s", err)
	}

	return nil
}

func(us *UserServ) GetUserByEmail(email string) (dto.UserProfileResponse, error) {
	user, err := us.userRepository.GetByEmail(email)
	if err != nil {
		return dto.UserProfileResponse{}, fmt.Errorf("get user profile by email %s", err)
	}

	resp := dto.UserProfileResponse(user)

	return resp, nil
}

func(us *UserServ) LoginUserByEmail(req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	user, err :=  us.userRepository.GetByEmail(req.Email)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	//compare hash password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.UserLoginResponse{}, fmt.Errorf("compare hashed password %s", err)
	}	

	//generate jwt token
	token, err := util.GenerateTokenJWT(int(user.Id), req.Email)
	if err != nil {
		return dto.UserLoginResponse{}, fmt.Errorf("generate jwt token %s", err)
	}

	resp := dto.UserLoginResponse{
		Token: token,
	}

	return resp, nil
}

func(us *UserServ) UpdateUserByEmail(req dto.UserUpdateProfileRequest) (dto.UserProfileResponse, error) {
	request := entity.User(req)
	user, err := us.userRepository.UpdateUserByEmail(&request)
	if err != nil {
		return dto.UserProfileResponse{}, fmt.Errorf("update user by email %s", err)
	}

	resp := dto.UserProfileResponse(user)

	return resp, nil
}

func(us *UserServ) UpdateUserImageByEmail(req dto.UserUpdateProfileRequest) (dto.UserProfileResponse, error) {
	request := entity.User(req)
	user, err := us.userRepository.UpdateImageByEmail(&request)
	if err != nil {
		return dto.UserProfileResponse{}, fmt.Errorf("update user image by email %s", err)
	}

	resp := dto.UserProfileResponse(user)

	return resp, nil
}