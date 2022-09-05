package responses

import "nidus-server/pkg/domain"

func UserSuccessResponse(data *domain.User, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
	}
}

func UsersSuccessResponse(data *[]domain.User, msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func UserErrorResponse(msg string) ResponseHTTP {
	return ResponseHTTP{
		Success: false,
		Message: msg,
		Data:    nil,
	}
}
