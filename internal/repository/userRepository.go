package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"secure-data-management/internal/model"
)

type UserRepository interface {
	FindById(ctx context.Context, userId string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Save(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
}

type MongoUserRepository struct {
	mongo          *mongo.Client
	userCollection *mongo.Collection
}

func (m *MongoUserRepository) FindById(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	filter := bson.M{"_id": userId}
	err := m.userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewMongoUserRepository(mongo *mongo.Client) *MongoUserRepository {
	userCollection := mongo.Database("file-upload").Collection("users")
	return &MongoUserRepository{mongo: mongo, userCollection: userCollection}
}

func (m *MongoUserRepository) FindByUsername(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	filter := bson.M{"username": id}
	err := m.userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MongoUserRepository) Update(ctx context.Context, user *model.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	_, err := m.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoUserRepository) Save(ctx context.Context, user *model.User) error {
	_, err := m.userCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
