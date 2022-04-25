package common_resp

type SuccessResponse struct {
	HttpReponse int         `json:"http_reponse"`
	RespCode    string      `json:"resp_code"`
	Status      string      `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

type ErrorResponse struct {
	HttpResponse int    `json:"http_response"`
	RespCode     string `json:"resp_code"`
	Status       string `json:"status"`
	Message      string `json:"message"`
}

var (
	responseSuccess = "00"
	responseFailed  = "02"
	statusSuccess   = "success"
	statusFailed    = "failed"
)

func SuccessMessage(message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		RespCode: responseSuccess,
		Status:   statusSuccess,
		Message:  message,
		Data:     data,
	}
}

func FailedMessage(message string) *ErrorResponse {
	return &ErrorResponse{
		RespCode: responseFailed,
		Status:   statusFailed,
		Message:  message,
	}
}
