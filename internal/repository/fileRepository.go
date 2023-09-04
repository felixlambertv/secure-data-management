package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"secure-data-management/internal/model"
)

type FileRepository interface {
	Save(context context.Context, file *model.File) error
	FindById(context context.Context, id string) (*model.File, error)
}

type MongoFileRepository struct {
	client         *mongo.Client
	fileCollection *mongo.Collection
}

func (m *MongoFileRepository) FindById(context context.Context, id string) (*model.File, error) {
	var file model.File
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	err := m.fileCollection.FindOne(context, filter).Decode(&file)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func NewMongoFileRepository(client *mongo.Client) *MongoFileRepository {
	fileCollection := client.Database("file-upload").Collection("files")
	return &MongoFileRepository{client: client, fileCollection: fileCollection}
}

func (m *MongoFileRepository) Save(context context.Context, file *model.File) error {
	_, err := m.fileCollection.InsertOne(context, file)
	if err != nil {
		return err
	}
	return nil
}
