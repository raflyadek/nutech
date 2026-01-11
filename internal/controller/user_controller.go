package controller

import (
	"fmt"
	"nutech-test/internal/dto"
	"nutech-test/util"
	"path/filepath"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	CreateUser(req dto.UserRegisterRequest) error
	GetUserProfileByEmail(email string) (dto.UserProfileResponse, error)
	LoginUserByEmail(req dto.UserLoginRequest) (dto.UserLoginResponse, error)
	UpdateUserByEmail(req dto.UserUpdateProfileRequest, email string) (dto.UserProfileResponse, error)
	UpdateUserImageByEmail(req dto.UserUpdateImageRequest, email string) (dto.UserProfileResponse, error)
	GetBalanceByEmail(email string) (dto.SaldoResponse, error)
	UpdateBalanceByEmail(req dto.TopUpSaldoRequest, email string) (dto.SaldoResponse, error)
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
		return util.UnauthorizedResponse(c, "Token tidak valid atau kadaluwarsa")
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
		return util.UnauthorizedResponse(c, "Token tidak valid atau kadaluwarsa")
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
		return util.UnauthorizedResponse(c, "Token tidak valid atau kadaluwarsa")
	}

	claim := user.Claims.(jwt.MapClaims)
	//get email from jwt payload
	email := claim["email"].(string)	

	//retrieve uploaded image/file
	file, err := c.FormFile("image")
	if err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	//validate image png/jpeg
	if err := util.ValidateImage(file); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	//path to save uploaded file
	uniqueId := uuid.New().String()
	fileName := uniqueId +file.Filename
	pathImage := filepath.Join("images", fileName)

	//save the file to the specified path
	if err := util.SaveUploadFile(file, pathImage); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	pictureUrl := fmt.Sprintf("http://localhost:8080/images/%s%s", uniqueId, file.Filename)

	req := dto.UserUpdateImageRequest{
		ProfileImage: pictureUrl,
	}

	resp, err := uc.userService.UpdateUserImageByEmail(req, email)
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Update Profile Image berhasil", resp)
}

func (uc *UserController) GetBalanceByEmail(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	//check if jwt exists safely
	if !ok || user == nil || !user.Valid {
		return util.UnauthorizedResponse(c, "Token tidak valid atau kadaluwarsa")
	}

	claim := user.Claims.(jwt.MapClaims)
	//get email from jwt payload
	email := claim["email"].(string)		

	resp, err := uc.userService.GetBalanceByEmail(email)
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Get Balance Berhasil", resp)
}

func (uc *UserController) UpdateBalanceByEmail(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	//check if jwt exists safely
	if !ok || user == nil || !user.Valid {
		return util.UnauthorizedResponse(c, "Token tidak valid atau kadaluwarsa")
	}

	claim := user.Claims.(jwt.MapClaims)
	//get email from jwt payload
	email := claim["email"].(string)

	//validate request
	req := new(dto.TopUpSaldoRequest)

	//bind from request
	if err := c.Bind(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	//validate request
	if err := uc.validate.Struct(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	resp, err := uc.userService.UpdateBalanceByEmail(*req, email)
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Top Up Balance Berhasil", resp)
}