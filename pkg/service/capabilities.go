package service

import (
	"fmt"
	"nidus-server/pkg/domain"
	"nidus-server/pkg/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CapabilityService interface {
	ListCapabilities() (*[]domain.Capability, error)
	CreateCapability(*domain.Capability) (*domain.Capability, error)
	ReadCapability(ID string) (*domain.Capability, error)
	UpdateCapability(capability *domain.Capability) (*domain.Capability, error)
	DeleteCapability(ID string) error
	CapabilityExist(capability *domain.Capability) (bool, error)
	VerifyDeviceCapabilities(capabilities []string) (*[]primitive.ObjectID, error)
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
	fmt.Println("ListCapabilities")
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
func (s *capabilityService) CapabilityExist(capability *domain.Capability) (bool, error) {
	// TODO to be implemented
	return true, nil
}

// The capabilities is a list of capability type and kind separated by a slash
// ie: 'sensor/temperature'
func (s *capabilityService) VerifyDeviceCapabilities(capabilities []string) (*[]primitive.ObjectID, error) {
	// Store valid capabilities
	validCapabilityIDs := []primitive.ObjectID{}

	// Loop through all capabilities
	for _, c := range capabilities {
		// TODO: Implement capability verification, with or without
		// the CapabilityExist method.
		fmt.Println(c)
	}

	return &validCapabilityIDs, nil

}
