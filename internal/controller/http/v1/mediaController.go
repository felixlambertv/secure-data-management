package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"secure-data-management/internal/controller/requests"
	"secure-data-management/internal/controller/responses"
	"secure-data-management/internal/middleware"
	"secure-data-management/internal/service"
)

type MediaController struct {
	fileService service.FileService
}

func NewMediaController(handler *gin.RouterGroup, fileService service.FileService) {
	r := &MediaController{fileService: fileService}
	h := handler.Group("/media").Use(middleware.TokenMiddleware)
	{
		h.POST("", r.UploadMedia)
		h.GET(":id", r.GetMedia)
	}
}

func (c *MediaController) GetMedia(context *gin.Context) {
	id := context.Param("id")
	token, exists := context.Get("token")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in context"})
		return
	}

	claims := token.(*jwt.Token).Claims.(jwt.MapClaims)

	media, err := c.fileService.GetMedia(id, claims["sub"].(string))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Something wrong"})
		return
	}

	context.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success get media",
		Status:  true,
		Data: gin.H{
			"url": media,
		},
	})
}

func (c *MediaController) UploadMedia(context *gin.Context) {
	token, exists := context.Get("token")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in context"})
		return
	}

	tokenString, exists := context.Get("tokenString")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found in context"})
		return
	}

	claims := token.(*jwt.Token).Claims.(jwt.MapClaims)

	var request requests.UploadFileRequest
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "request invalid",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	err := c.fileService.UploadFile(request, claims["sub"].(string), tokenString.(string))
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.ErrorResponse{
			Message: "fail upload",
			Debug:   err,
			Errors:  err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, responses.SuccessResponse{
		Message: "Success upload file",
		Status:  true,
		Data:    nil,
	})
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
