package service

import (
	"nidus-server/pkg/domain"
	"nidus-server/pkg/repository"
)

type ZoneService interface {
	ListZones() (*[]domain.Zone, error)
	CreateZone(*domain.Zone) (*domain.Zone, error)
	ReadZone(ID string) (*domain.Zone, error)
	UpdateZone(user *domain.Zone) (*domain.Zone, error)
	DeleteZone(ID string) error
}

type zoneService struct {
	repository repository.ZoneRepository
}

func NewZoneService(r repository.ZoneRepository) ZoneService {
	return &zoneService{
		repository: r,
	}
}

func (s *zoneService) ListZones() (*[]domain.Zone, error) {
	return s.repository.ListZones()
}

func (s *zoneService) CreateZone(user *domain.Zone) (*domain.Zone, error) {
	return s.repository.CreateZone(user)
}

func (s *zoneService) ReadZone(ID string) (*domain.Zone, error) {
	return s.repository.ReadZone(ID)
}

func (s *zoneService) UpdateZone(user *domain.Zone) (*domain.Zone, error) {
	return s.repository.UpdateZone(user)
}

func (s *zoneService) DeleteZone(ID string) error {
	return s.repository.DeleteZone(ID)
}
