package responses

import "nidus-server/pkg/domain"

func ListMeasurementSuccessResponse(data *[]domain.Measurement, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func CreateMeasurementSuccessResponse(data *domain.Measurement, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func ReadMeasurementSuccessResponse(data *domain.Measurement, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func MeasurementSuccessResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    nil,
	}
}

func MeasurementErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: msg,
		Data:    nil,
	}
}
