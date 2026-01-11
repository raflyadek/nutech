package repository

import (
	"context"
	"database/sql"
	"nutech-test/internal/entity"
)

type ServiceRepo struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) *ServiceRepo {
	return &ServiceRepo{db: db}
}

func (sr *ServiceRepo) GetAllServices() ([]entity.Service, error) {
	rows, err := sr.db.QueryContext(context.Background(), `
	SELECT service_code, service_name, service_icon, service_tariff FROM service`)
	if err != nil {
		return []entity.Service{}, err
	}
	defer rows.Close()

	var service []entity.Service

	for rows.Next() {
		var s entity.Service
		if err := rows.Scan(&s.ServiceCode, &s.ServiceName, &s.ServiceIcon, &s.ServiceTariff); err != nil {
			return []entity.Service{}, err
		}
		
		service = append(service, s)
	}

	return service, nil
}