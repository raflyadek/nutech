package controller

import (
	"nutech-test/internal/dto"
	"nutech-test/util"

	"github.com/labstack/echo/v4"
)

type ServiceService interface {
	GetAllService() ([]dto.Service, error)
}

type ServiceController struct {
	serviceService ServiceService
}

func NewServiceController(ss ServiceService) *ServiceController {
	return &ServiceController{serviceService: ss}
}

func (sc *ServiceController) GetAllService(c echo.Context) error {
	resp, err := sc.serviceService.GetAllService()
	if err != nil {
		return util.InternalServerErrorResponse(c, err.Error())
	}

	return util.SuccessResponse(c, "Sukses", resp)
}