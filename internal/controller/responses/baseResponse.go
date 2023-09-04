package responses

type ErrorResponse struct {
	Message string `json:"message"`
	Debug   error  `json:"debug"`
	Errors  any    `json:"errors"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Data    any    `json:"data"`
}
