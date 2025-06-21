package repositories

import (
	"backend/core/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUser(email string) (result models.UserModel, err error)

	CreateUser(user models.CreateUserModel) error

	UpdateUser(id string, user models.UpdateUserModel) (result models.UserModel, err error)
}
type userRepo struct {
	db         *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) UserRepository {
	return &userRepo{
		db:         db,
		collection: collection,
	}
}

func (r *userRepo) GetUser(email string) (result models.UserModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "email", Value: email}}
	err = r.db.Collection(r.collection).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *userRepo) CreateUser(user models.CreateUserModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.db.Collection(r.collection).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) UpdateUser(id string, user models.UpdateUserModel) (result models.UserModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "id", Value: id}}
	_, err = r.db.Collection(r.collection).UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: user}})
	if err != nil {
		return result, err
	}
	err = r.db.Collection(r.collection).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
