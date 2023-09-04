package responses

type ErrorRes struct {
	Message string
	Debug   error
	Errors  any
}
