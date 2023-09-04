package v1

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"secure-data-management/config"
	"secure-data-management/internal/repository"
	"secure-data-management/internal/service"
)

func NewRouter(handler *gin.Engine, sess *session.Session, config *config.AWSConfig, db *mongo.Client) {
	handler.Use(gin.Recovery())

	userRepo := repository.NewMongoUserRepository(db)
	authService := service.NewCognitoAuthService(userRepo, sess, config)
	fileRepository := repository.NewMongoFileRepository(db)
	fileService := service.NewS3FileService(userRepo, fileRepository, sess)

	handler.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	h := handler.Group("api/v1")
	{
		NewAuthController(h, authService)
		NewMediaController(h, fileService)
	}
}
