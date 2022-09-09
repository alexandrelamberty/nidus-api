package responses

import "nidus-server/pkg/domain"

func MeasurementsuccessResponse(data *domain.Measurement, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
	}
}

func MeasurementsSuccessResponse(data *[]domain.Measurement, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func MeasurementErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: msg,
		Data:    nil,
	}
}
