package repository

import (
	"context"
	"database/sql"
	"nutech-test/internal/entity"
)

type BannerRepo struct {
	db *sql.DB
}

func NewBannerRepository(db *sql.DB) *BannerRepo {
	return &BannerRepo{db: db}
}

func (br *BannerRepo) GetAllBanner() ([]entity.Banner, error) {
	rows, err := br.db.QueryContext(context.Background(), `
	SELECT banner_name, banner_image, description FROM banner`)
	if err != nil {
		return []entity.Banner{}, err
	}
	defer rows.Close()

	var banner []entity.Banner

	for rows.Next() {
		var b entity.Banner
		if err := rows.Scan(&b.BannerName, &b.BannerImage, &b.Description); err != nil {
			return []entity.Banner{}, err
		}
		
		banner = append(banner, b)
	}

	return banner, nil
}
