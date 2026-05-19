package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BuildResponseSuccess(message string, data any) Response {

	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func BuildResponseFailed(message string) Response {
	return Response{
		Status:  false,
		Message: message,
	}
}
