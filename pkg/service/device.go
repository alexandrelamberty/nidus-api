package service

import (
	"nidus-server/pkg/domain"
	"nidus-server/pkg/repository"
)

type DeviceService interface {
	ListDevices() (*[]domain.Device, error)
	CreateDevice(*domain.Device) (*domain.Device, error)
	ReadDevice(id string) (*domain.Device, error)
	UpdateDevice(user *domain.Device) (*domain.Device, error)
	DeleteDevice(id string) error
	PairDevice(id string) string
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

func (s *deviceService) CreateDevice(device *domain.Device) (*domain.Device, error) {
	return s.repository.CreateDevice(device)
}

func (s *deviceService) ReadDevice(ID string) (*domain.Device, error) {
	return s.repository.ReadDevice(ID)
}

func (s *deviceService) UpdateDevice(device *domain.Device) (*domain.Device, error) {
	return s.repository.UpdateDevice(device)
}

func (s *deviceService) DeleteDevice(id string) error {
	return s.repository.DeleteDevice(id)
}

func (s *deviceService) PairDevice(id string) string {
	return "pairing: " + id
}
