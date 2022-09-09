package service

import (
	"nidus-server/pkg/domain"
	"nidus-server/pkg/repository"
)

type MeasurementService interface {
	ListMeasurements() (*[]domain.Measurement, error)
	CreateMeasurement(*domain.Measurement) (*domain.Measurement, error)
	ReadMeasurement(ID string) (*domain.Measurement, error)
	UpdateMeasurement(user *domain.Measurement) (*domain.Measurement, error)
	DeleteMeasurement(ID string) error
}

type measurementService struct {
	repository repository.MeasurementRepository
}

func NewMeasurementService(r repository.MeasurementRepository) MeasurementService {
	return &measurementService{
		repository: r,
	}
}

func (s *measurementService) ListMeasurements() (*[]domain.Measurement, error) {
	return s.repository.ListMeasurements()
}

func (s *measurementService) CreateMeasurement(user *domain.Measurement) (*domain.Measurement, error) {
	return s.repository.CreateMeasurement(user)
}

func (s *measurementService) ReadMeasurement(ID string) (*domain.Measurement, error) {
	return s.repository.ReadMeasurement(ID)
}

func (s *measurementService) UpdateMeasurement(user *domain.Measurement) (*domain.Measurement, error) {
	return s.repository.UpdateMeasurement(user)
}

func (s *measurementService) DeleteMeasurement(ID string) error {
	return s.repository.DeleteMeasurement(ID)
}
