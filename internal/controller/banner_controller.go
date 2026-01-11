package controller

import (
	"nutech-test/internal/dto"
	"nutech-test/util"

	"github.com/labstack/echo/v4"
)

type BannerService interface {
	GetAllBanner() ([]dto.BannerResponse, error)
}

type BannerController struct {
	bannerService BannerService
}

func NewBannerController(bs BannerService) *BannerController {
	return &BannerController{bannerService: bs}
}

func (bc *BannerController) GetAllBanner(c echo.Context) error {
	resp, err := bc.bannerService.GetAllBanner()
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Sukses", resp)
}