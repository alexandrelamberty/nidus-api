package service

import (
	"nidus-server/pkg/domain"
	"nidus-server/pkg/repository"
)

type DeviceService interface {
	ListDevices() (*[]domain.Device, error)
	CreateDevice(*domain.Device) (*domain.Device, error)
	ReadDevice(ID string) (*domain.Device, error)
	UpdateDevice(user *domain.Device) (*domain.Device, error)
	DeleteDevice(ID string) error
}

type deviceService struct {
	repository repository.DeviceRepository
}

func NewDeviceService(r repository.DeviceRepository) DeviceService {
	return &deviceService{
		repository: r,
	}
}

func (s *deviceService) ListDevices() (*[]domain.Device, error) {
	return s.repository.ListDevices()
}

func (s *deviceService) CreateDevice(user *domain.Device) (*domain.Device, error) {
	return s.repository.CreateDevice(user)
}

func (s *deviceService) ReadDevice(ID string) (*domain.Device, error) {
	return s.repository.ReadDevice(ID)
}

func (s *deviceService) UpdateDevice(user *domain.Device) (*domain.Device, error) {
	return s.repository.UpdateDevice(user)
}

func (s *deviceService) DeleteDevice(ID string) error {
	return s.repository.DeleteDevice(ID)
}
