/***
* =======================================================================
* FILE NAME:        tags_repository.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package repository

import (
	"context"
	"golang-mongodb/model/domain"
	"golang-mongodb/model/web"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagsRepository interface {
	Save(ctx context.Context, client *mongo.Client, tags web.TagsRequest) domain.Tags
	Update(ctx context.Context, client *mongo.Client, filter bson.M, tags web.TagsRequest) bool
	Delete(ctx context.Context, client *mongo.Client, selector bson.M) bool
	FindAll(ctx context.Context, client *mongo.Client) []domain.Tags
}
