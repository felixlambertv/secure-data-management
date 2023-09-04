package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"secure-data-management/config"
	v1 "secure-data-management/internal/controller/http/v1"
	"time"
)

func main() {
	fmt.Println(os.Environ())
	awsConfig := config.NewAWSConfig(
		"ap-southeast-1",
		"AKIAZTYI65C3DPQTNGM2",
		"X9LyJ+0LMU8M1AaJ8tlnopE+uf6QkQ1aJDMbFSzT",
		"ap-southeast-1_zFFMzwFGK",
		"3u2lo7kgtuf83pf4nfiouqgcp8",
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
