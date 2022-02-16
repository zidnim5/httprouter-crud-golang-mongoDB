/***
* =======================================================================
* FILE NAME:        news_repository.go
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

type NewsRepository interface {
	Save(ctx context.Context, client *mongo.Client, request web.NewsCreateRequest) domain.News
	Update(ctx context.Context, client *mongo.Client, filter bson.M, request bson.M) bool
	Delete(ctx context.Context, client *mongo.Client, filter bson.M, request bson.M) bool
	FindAll(ctx context.Context, client *mongo.Client, selector web.NewsFindRequest) []domain.News
}
