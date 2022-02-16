/***
* =======================================================================
* FILE NAME:        tags_repository_impl.go
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

type TagsRepositoryImpl struct {
}

/**
 * Init
 */
func NewTagsRepository() TagsRepository {
	return &TagsRepositoryImpl{}
}

func (c *TagsRepositoryImpl) Save(ctx context.Context, client *mongo.Client, tags web.TagsRequest) domain.Tags {
	collection := client.Database("golang-mongodb").Collection("tags")
	insertResult, err := collection.InsertOne(ctx, tags)
	if err != nil {
		helper.PanicIfError(err)
	}

	return domain.Tags{
		Id:   insertResult.InsertedID.(primitive.ObjectID).Hex(),
		Name: tags.Name,
	}
}

func (c *TagsRepositoryImpl) Update(ctx context.Context, client *mongo.Client, filter bson.M, tags web.TagsRequest) bool {
	updatedResult, err := client.Database("golang-mongodb").Collection("tags").UpdateOne(ctx, filter, bson.M{"$set": tags})
	if err != nil {
		helper.PanicIfError(err)
	}

	if updatedResult.MatchedCount > 0 {
		return true
	}

	return false
}

func (c *TagsRepositoryImpl) Delete(ctx context.Context, client *mongo.Client, selector bson.M) bool {
	deletedResult, err := client.Database("golang-mongodb").Collection("tags").DeleteOne(ctx, selector)
	if err != nil {
		helper.PanicIfError(err)
	}

	if deletedResult.DeletedCount > 0 {
		return true
	}
	return false
}

func (c *TagsRepositoryImpl) FindAll(ctx context.Context, client *mongo.Client) []domain.Tags {
	data := []domain.Tags{}

	findAllResult, _ := client.Database("golang-mongodb").Collection("tags").Find(ctx, bson.D{})

	for findAllResult.Next(ctx) {
		var tmp domain.Tags
		err := findAllResult.Decode(&tmp)
		helper.PanicIfError(err)

		data = append(data, tmp)
	}

	return data
}
