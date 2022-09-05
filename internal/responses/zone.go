package responses

import "nidus-server/pkg/domain"

func ZoneSuccessResponse(data *domain.Zone, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
	}
}

func ZonesSuccessResponse(data *[]domain.Zone, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func ZoneErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: msg,
		Data:    nil,
	}
}
