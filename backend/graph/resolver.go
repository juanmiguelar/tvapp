package graph

import (
	"context"
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
		ID:        primitive.NewObjectID().Hex(),
		Title:     title,
		Content:   content,
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
