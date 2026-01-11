package controller

import (
	"nutech-test/internal/dto"
	"nutech-test/util"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TransactionService interface {
	CreateTransaction(req dto.TransactionRequest, email string) (dto.TransactionResponse, error)
	GetAllTransactionByEmail(email string) ([]dto.TransactionHistoryResponse, error)
}

type TransactionController struct {
	transactionService TransactionService
	validate *validator.Validate
}

func NewTransactionController(ts TransactionService, validate *validator.Validate) *TransactionController {
	return &TransactionController{transactionService: ts, validate: validate}
}

func (tc *TransactionController) CreateTransaction(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	//check if jwt exists safely
	if !ok || user == nil || !user.Valid {
		return util.UnauthorizedResponse(c, "Token tidak valid atau kadaluwarsa")
	}

	claim := user.Claims.(jwt.MapClaims)
	//get email from jwt payload
	email := claim["email"].(string)

	req := new(dto.TransactionRequest)

	//bind from request
	if err := c.Bind(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	//validate request
	if err := tc.validate.Struct(req); err != nil {
		return util.BadRequestResponse(c, err.Error())
	}

	resp, err := tc.transactionService.CreateTransaction(*req, email)
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Transaksi berhasil", resp)
}

func (tc *TransactionController) GetAllTransactionByEmail(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	//check if jwt exists safely
	if !ok || user == nil || !user.Valid {
		return util.UnauthorizedResponse(c, "Token tidak valid atau kadaluwarsa")
	}

	claim := user.Claims.(jwt.MapClaims)
	//get email from jwt payload
	email := claim["email"].(string)	

	resp, err := tc.transactionService.GetAllTransactionByEmail(email)
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Get History Berhasil", resp)
}
