package responses

import "nidus-server/pkg/domain"

func DeviceSuccessResponse(data *domain.Device, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Data:    data,
		Message: msg,
	}
}

func DevicesSuccessResponse(data *[]domain.Device, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func DeviceErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: msg,
		Data:    nil,
	}
}
