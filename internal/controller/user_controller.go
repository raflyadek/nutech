package controller

import (
	"nutech-test/internal/dto"
	"nutech-test/util"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

type UserService interface {
	CreateUser(req dto.UserRegisterRequest) error
	GetUserProfileByEmail(email string) (dto.UserProfileResponse, error)
	LoginUserByEmail(req dto.UserLoginRequest) (dto.UserLoginResponse, error)
	UpdateUserByEmail(req dto.UserUpdateProfileRequest, email string) (dto.UserProfileResponse, error)
	UpdateUserImageByEmail(req dto.UserUpdateImageRequest, email string) (dto.UserProfileResponse, error)
}

type UserController struct {
	userService UserService
	validate *validator.Validate
}

func NewUserController(us UserService, validate *validator.Validate) *UserController {
	return &UserController{userService: us, validate: validate}
}

func (uc *UserController) CreateUser(c echo.Context) error {
	//create instance from struct
	req := new(dto.UserRegisterRequest)

	//bind from request
	if err := c.Bind(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	//validate request
	if err := uc.validate.Struct(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	if err := uc.userService.CreateUser(*req); err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Registrasi berhasil silahkan login", map[string]interface{}{})
}

func (uc *UserController) LoginUser(c echo.Context) error {
	req := new(dto.UserLoginRequest)

	//bind from request
	if err := c.Bind(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	//validate request
	if err := uc.validate.Struct(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	resp, err := uc.userService.LoginUserByEmail(*req)
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Login sukses", resp)
}

func (uc *UserController) GetUserProfileByEmail(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	//check safely jwt exists safely
	if !ok || user == nil || !user.Valid {
		return util.UnauthorizedResponse(c, "Token tidak tidak valid atau kadaluwarsa")
	}

	claim := user.Claims.(jwt.MapClaims)
	//get email from jwt payload
	email := claim["email"].(string)

	resp, err := uc.userService.GetUserProfileByEmail(email)
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Sukses", resp)
}

func (uc *UserController) UpdateUserByEmail(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	//check if jwt exists safely
	if !ok || user == nil || !user.Valid {
		return util.UnauthorizedResponse(c, "Token tidak tidak valid atau kadaluwarsa")
	}

	claim := user.Claims.(jwt.MapClaims)
	//get email from jwt payload
	email := claim["email"].(string)	

	req := new(dto.UserUpdateProfileRequest)

	//bind from request
	if err := c.Bind(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	//validate request
	if err := uc.validate.Struct(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	resp, err := uc.userService.UpdateUserByEmail(*req, email)
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Update profile berhasil", resp)
}

func (uc *UserController) UpdateUserImageByEmail(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	//check if jwt exists safely
	if !ok || user == nil || !user.Valid {
		return util.UnauthorizedResponse(c, "Token tidak tidak valid atau kadaluwarsa")
	}

	claim := user.Claims.(jwt.MapClaims)
	//get email from jwt payload
	email := claim["email"].(string)	

	req := new(dto.UserUpdateImageRequest)

	//bind from request
	if err := c.Bind(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	//validate request
	if err := uc.validate.Struct(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}	

	resp, err := uc.userService.UpdateUserImageByEmail(*req, email)
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Update Profile Image berhasil", resp)
}
