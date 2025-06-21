package repositories

import (
	"backend/core/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuoteRepository interface {
	GetQuotes() (result []models.QuoteModel, err error)

	GetQuote(id string) (result models.QuoteModel, err error)

	CreateQuote(quote models.RepoCreateQuoteModel) (result models.QuoteModel, err error)

	UpdateQuote(id string, payload models.RepoUpdateQuoteModel) (result models.QuoteModel, err error)

	DeleteQuote(id string) error
}

type QuoteRepo struct {
	db         *mongo.Database
	collection string
}

func NewQuoteRepository(db *mongo.Database, collection string) QuoteRepository {
	return &QuoteRepo{
		db:         db,
		collection: collection,
	}
}

func (r *QuoteRepo) GetQuotes() (result []models.QuoteModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.D{})
	if err != nil {
		return result, err
	}
	if err = cursor.All(ctx, &result); err != nil {
		return result, err
	}

	return result, nil
}

func (r *QuoteRepo) GetQuote(id string) (result models.QuoteModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "id", Value: id}}
	err = r.db.Collection(r.collection).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *QuoteRepo) CreateQuote(payload models.RepoCreateQuoteModel) (result models.QuoteModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = r.db.Collection(r.collection).InsertOne(ctx, payload)
	if err != nil {
		return result, err
	}
	err = r.db.Collection(r.collection).FindOne(ctx, bson.D{{Key: "id", Value: payload.ID}}).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *QuoteRepo) UpdateQuote(id string, payload models.RepoUpdateQuoteModel) (result models.QuoteModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "id", Value: id}}

	err = r.db.Collection(r.collection).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	_, err = r.db.Collection(r.collection).UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: payload}})
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *QuoteRepo) DeleteQuote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "id", Value: id}}
	_, err := r.db.Collection(r.collection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
