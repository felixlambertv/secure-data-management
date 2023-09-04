package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"secure-data-management/config"
	v1 "secure-data-management/internal/controller/http/v1"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	awsConfig := config.NewAWSConfig(
		os.Getenv("AWS_REGION"),
		os.Getenv("AWS_ACCESS_KEY_ID"),
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		os.Getenv("COGNITO_USER_POOL_ID"),
		os.Getenv("COGNITO_CLIENT_ID"),
	)

	connectionString := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)
	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic("Failed connect to db: " + err.Error())
	}

	sess, err := awsConfig.NewSession()
	if err != nil {
		panic("Failed start aws session: " + err.Error())
	}

	handler := gin.Default()
	v1.NewRouter(handler, sess, awsConfig, db)
	err = handler.Run(":8080")
	if err != nil {
		panic("Failed run server: " + err.Error())
	}
}
