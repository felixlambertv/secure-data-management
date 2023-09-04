package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"secure-data-management/internal/controller/requests"
	"secure-data-management/internal/controller/responses"
	"secure-data-management/internal/service"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(handler *gin.RouterGroup, authService service.AuthService) {
	r := &AuthController{authService: authService}
	h := handler.Group("users")
	{
		h.POST("/login", r.login)
		h.POST("/register", r.register)
		h.POST("/verify", r.verify)
	}
}

func (c *AuthController) register(context *gin.Context) {
	var request requests.UserRegisterRequest
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorRes{
			Message: "request invalid",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	err := c.authService.Register(request)
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorRes{
			Message: "Fail user register",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true})
}

func (c *AuthController) verify(context *gin.Context) {
	var request requests.AccountVerifyRequest
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorRes{
			Message: "request invalid",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	err := c.authService.Verify(request)
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorRes{
			Message: "Fail verify user",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true})
}

func (c *AuthController) login(context *gin.Context) {
	var request requests.LoginRequest
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorRes{
			Message: "request invalid",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	data, err := c.authService.Login(request)
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorRes{
			Message: "Fail verify user",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, data)
}
