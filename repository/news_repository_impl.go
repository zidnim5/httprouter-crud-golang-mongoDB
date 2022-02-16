/***
* =======================================================================
* FILE NAME:        news_repository_impl.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package repository

import (
	"context"
	"golang-mongodb/helper"
	"golang-mongodb/model/domain"
	"golang-mongodb/model/web"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsRepositoryImpl struct {
}

/**
 * Init
 */
func NewNewsRepository() NewsRepository {
	return &NewsRepositoryImpl{}
}

func (c *NewsRepositoryImpl) Save(ctx context.Context, client *mongo.Client, request web.NewsCreateRequest) domain.News {
	insertResult, err := client.Database("golang-mongodb").Collection("news").InsertOne(ctx, request)
	if err != nil {
		helper.PanicIfError(err)
	}

	return domain.News{
		Id:      insertResult.InsertedID.(primitive.ObjectID).Hex(),
		Topic:   request.Topic,
		Title:   request.Title,
		Status:  request.Status,
		Tags:    request.Tags,
		Content: request.Content,
	}
}

func (c *NewsRepositoryImpl) Update(ctx context.Context, client *mongo.Client, filter bson.M, request bson.M) bool {
	updatedResult, err := client.Database("golang-mongodb").Collection("news").UpdateOne(ctx, filter, bson.M{"$set": request})
	if err != nil {
		helper.PanicIfError(err)
	}

	if updatedResult.MatchedCount > 0 {
		return true
	}
	return false
}

func (c *NewsRepositoryImpl) Delete(ctx context.Context, client *mongo.Client, filter bson.M, request bson.M) bool {
	updatedResult, err := client.Database("golang-mongodb").Collection("news").UpdateOne(ctx, filter, bson.M{"$set": request})
	if err != nil {
		helper.PanicIfError(err)
	}

	if updatedResult.MatchedCount > 0 {
		return true
	}

	return false
}

func (c *NewsRepositoryImpl) FindAll(ctx context.Context, client *mongo.Client, selector web.NewsFindRequest) []domain.News {
	check := web.NewsFindRequest{}
	var filter bson.M

	switch {
	case selector != check:
		// todo: find with parameter ? topic / status
		filter = bson.M{
			"topic": bson.M{
				"$regex": primitive.Regex{
					Pattern: selector.Topic,
					Options: "i",
				},
			},
			"$or": []bson.M{
				{
					"status": bson.M{
						"$regex": primitive.Regex{
							Pattern: selector.Status,
							Options: "i",
						},
					},
				},
			},
		}
	default:
		// todo: find without parameter
		filter = bson.M{}
	}

	data := []domain.News{}

	findAllResult, _ := client.Database("golang-mongodb").Collection("news").Find(ctx, filter)

	for findAllResult.Next(ctx) {
		var tmp domain.News
		err := findAllResult.Decode(&tmp)
		helper.PanicIfError(err)

		data = append(data, tmp)
	}

	return data
}
