package service

import (
	"nidus-server/internal/requests"
	"nidus-server/pkg/domain"
	"nidus-server/pkg/repository"
)

type MeasurementService interface {
	ListMeasurements() (*[]domain.Measurement, error)
	CreateMeasurement(*requests.CreateMeasurementRequest) (*domain.Measurement, error)
	ReadMeasurement(deviceId string, sensorType string, timestamp string) (*domain.Measurement, error)
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

func (s *measurementService) CreateMeasurement(measurement *requests.CreateMeasurementRequest) (*domain.Measurement, error) {
	return s.repository.CreateMeasurement(measurement)
}

func (s *measurementService) ReadMeasurement(ID string, sensorType string, timestamp string) (*domain.Measurement, error) {
	return s.repository.ReadMeasurement(ID, sensorType, timestamp)
}
