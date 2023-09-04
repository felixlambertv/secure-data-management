package requests

type UserRegisterRequest struct {
	Email    string
	Username string
	Password string
}

type AccountVerifyRequest struct {
	Username string
	Code     string
}

type LoginRequest struct {
	Username string
	Password string
}
