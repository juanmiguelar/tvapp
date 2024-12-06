package graph

import (
	"context"
	"errors"
	"tvapp-backend/database"
	"tvapp-backend/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Resolver struct {
	NewsCollection *mongo.Collection
}

func NewResolver() *Resolver {
	return &Resolver{
		NewsCollection: database.Client.Database("tvapp_db").Collection("news"),
	}
}

// Query: Get all news
func (r *Resolver) GetNews(ctx context.Context) ([]*model.News, error) {
	cursor, err := r.NewsCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var news []*model.News
	for cursor.Next(ctx) {
		var n model.News
		if err := cursor.Decode(&n); err != nil {
			return nil, err
		}
		news = append(news, &n)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return news, nil
}

// Mutation: Create a news item
func (r *Resolver) CreateNews(ctx context.Context, title string, content string, authorName string, authorEmail string) (*model.News, error) {
	news := model.News{
		ID:      primitive.NewObjectID().Hex(),
		Title:   title,
		Content: content,
		Author: &model.Author{
			Name:  authorName,
			Email: authorEmail,
		},
	}

	_, err := r.NewsCollection.InsertOne(ctx, news)
	if err != nil {
		return nil, err
	}
	return &news, nil
}

// Mutation: Update a news item
func (r *Resolver) UpdateNews(ctx context.Context, id string, title *string, content *string, authorName *string, authorEmail *string) (*model.News, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	update := bson.M{}
	if title != nil {
		update["title"] = *title
	}
	if content != nil {
		update["content"] = *content
	}
	if authorName != nil || authorEmail != nil {
		author := bson.M{}
		if authorName != nil {
			author["name"] = *authorName
		}
		if authorEmail != nil {
			author["email"] = *authorEmail
		}
		update["author"] = author
	}

	result := r.NewsCollection.FindOneAndUpdate(ctx, bson.M{"_id": objectID}, bson.M{"$set": update}, nil)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updatedNews model.News
	if err := result.Decode(&updatedNews); err != nil {
		return nil, err
	}

	return &updatedNews, nil
}

// Mutation: Delete a news item
func (r *Resolver) DeleteNews(ctx context.Context, id string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, errors.New("invalid ID format")
	}

	result, err := r.NewsCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return false, err
	}

	return result.DeletedCount > 0, nil
}
