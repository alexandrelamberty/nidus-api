package responses

import "nidus-server/pkg/domain"

func CapabilitySuccessResponse(data *domain.Capability, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
	}
}

func CapabilitiesSuccessResponse(data *[]domain.Capability, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func CapabilityErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: msg,
		Data:    nil,
	}
}
