package service

import (
	"fmt"
	"nutech-test/internal/dto"
	"nutech-test/internal/entity"
)

type BannerRepository interface {
	GetAllBanner() ([]entity.Banner, error)
}

type BannerServ struct {
	bannerRepository BannerRepository
}

func NewBannerService(br BannerRepository) *BannerServ {
	return &BannerServ{bannerRepository: br}
}

func (bs *BannerServ) GetAllBanner() ([]dto.BannerResponse, error) {
	banners, err := bs.bannerRepository.GetAllBanner()
	if err != nil {
		return []dto.BannerResponse{}, fmt.Errorf("get all banner %s", err)
	}

	var resp []dto.BannerResponse

	for _, banner := range banners {
		resp = append(resp, dto.BannerResponse(banner))
	}

	return resp, nil
}
