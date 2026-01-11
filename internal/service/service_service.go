package service

import (
	"fmt"
	"nutech-test/internal/dto"
	"nutech-test/internal/entity"
)

type ServiceRepository interface {
	GetAllServices() ([]entity.Service, error)
}

type ServiceServ struct {
	serviceRepository ServiceRepository
}

func NewServiceService(sr ServiceRepository) *ServiceServ {
	return &ServiceServ{serviceRepository: sr}
}

func (ss *ServiceServ) GetAllService() ([]dto.Service, error) {
	services, err := ss.serviceRepository.GetAllServices()
	if err != nil {
		return []dto.Service{}, fmt.Errorf("get all service %s", err)
	}

	var resp []dto.Service
	for _, service := range services {
		resp = append(resp, dto.Service(service))
	}

	return resp, nil
}