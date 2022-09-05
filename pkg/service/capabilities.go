package service

import (
	"nidus-server/pkg/domain"
	"nidus-server/pkg/repository"
)

type CapabilityService interface {
	ListCapabilities() (*[]domain.Capability, error)
	CreateCapability(*domain.Capability) (*domain.Capability, error)
	ReadCapability(ID string) (*domain.Capability, error)
	UpdateCapability(capability *domain.Capability) (*domain.Capability, error)
	DeleteCapability(ID string) error
}

type capabilityService struct {
	repository repository.CapabilityRepository
}

func NewCapabilityService(r repository.CapabilityRepository) CapabilityService {
	return &capabilityService{
		repository: r,
	}
}

func (s *capabilityService) ListCapabilities() (*[]domain.Capability, error) {
	return s.repository.ListCapabilities()
}

func (s *capabilityService) CreateCapability(capability *domain.Capability) (*domain.Capability, error) {
	return s.repository.CreateCapability(capability)
}

func (s *capabilityService) ReadCapability(ID string) (*domain.Capability, error) {
	return s.repository.ReadCapability(ID)
}

func (s *capabilityService) UpdateCapability(capability *domain.Capability) (*domain.Capability, error) {
	return s.repository.UpdateCapability(capability)
}

func (s *capabilityService) DeleteCapability(ID string) error {
	return s.repository.DeleteCapability(ID)
}
