package helper

type Response struct {
	StatusCode int         `json:"status_code"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func APIResponse(message string, code int, status bool, data interface{}) Response {
	jsonResponse := Response{
		StatusCode: code,
		Success:    status,
		Message:    message,
		Data:       data,
	}

	return jsonResponse
}
