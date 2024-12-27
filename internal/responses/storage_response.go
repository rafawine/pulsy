package responses

type Respose struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(message string, data interface{}) Respose {
	return Respose{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func Error(message string, errors interface{}) Respose {
	return Respose{
		Status:  "error",
		Message: message,
		Error:   errors,
	}
}
