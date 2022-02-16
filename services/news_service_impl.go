/***
* =======================================================================
* FILE NAME:        news_service_impl.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package services

import (
	"context"
	"fmt"
	"golang-mongodb/exception"
	"golang-mongodb/helper"
	"golang-mongodb/model/web"
	"golang-mongodb/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsServiceImpl struct {
	NewsRepository repository.NewsRepository
	DB             *mongo.Client
	Rds            *redis.Client
	validate       *validator.Validate
}

var ttl time.Duration = time.Duration(5) * time.Minute

/**
 * Init
 */
func NewNewsService(newsRepository repository.NewsRepository, DB *mongo.Client, Rds *redis.Client, validate *validator.Validate) NewsService {
	return &NewsServiceImpl{NewsRepository: newsRepository, DB: DB, Rds: Rds, validate: validate}
}

/**
 * todo: create new news
 *
 * @param context.Context
 * @param web.newsRequest
 *
 * @return web.newsResponse
 */
func (c *NewsServiceImpl) Create(ctx context.Context, request web.NewsCreateRequest) web.NewsResponse {
	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	createdResult := c.NewsRepository.Save(ctx, c.DB, request)

	result := web.NewsResponse(createdResult)

	return result

}

/**
 * todo: update news
 *
 * @param context.Context
 * @param web.newsUpdateRequest
 *
 * @return web.newsResponse
 */
func (c *NewsServiceImpl) Update(ctx context.Context, request web.NewsUpdateRequest) web.NewsResponse {
	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	objectIDS, _ := primitive.ObjectIDFromHex(request.Id)
	filter := bson.M{"_id": objectIDS}
	changes := bson.M{
		"topic":   request.Topic,
		"title":   request.Title,
		"status":  request.Status,
		"tags":    request.Tags,
		"content": request.Content,
	}

	updatedResult := c.NewsRepository.Update(ctx, c.DB, filter, changes)
	if updatedResult {
		return web.NewsResponse(request)
	}

	panic(exception.NewNotFoundError("Data not found"))
}

/**
 * todo: delete news
 *
 * @param context.Context
 * @param web.newsDeleteRequest
 *
 */
func (c *NewsServiceImpl) Delete(ctx context.Context, request web.NewsDeleteRequest) interface{} {
	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	objectIDS, _ := primitive.ObjectIDFromHex(request.Id)
	selector := bson.M{"_id": objectIDS}
	change := bson.M{"status": "deleted"}

	deletedResult := c.NewsRepository.Delete(ctx, c.DB, selector, change)
	if deletedResult {
		return nil
	}

	panic(exception.NewNotFoundError("Data not found"))
}

/**
 * todo: findall news
 *
 * @param context.Context
 *
 * @return []domain.news
 *
 */
func (c *NewsServiceImpl) FindAll(ctx context.Context, request web.NewsFindRequest) interface{} {
	findRequest := web.NewsFindRequest{}
	var findAllResult interface{}
	var key string

	if request != findRequest {

		// get to redis
		if request.Status != "" {
			fmt.Println("Get from redis")
			logrus.Info("Get from redis")
			findAllResult = helper.GetRedis(c.Rds, request.Status)
			key = request.Status
		} else if request.Topic != "" {
			fmt.Println("Get from redis")
			logrus.Info("Get from redis")
			findAllResult = helper.GetRedis(c.Rds, request.Topic)
			key = request.Topic
		}

		// convert interface{} to struct{}
		if findAllResult == nil || findAllResult == "" {
			fmt.Println("Get from DB")
			logrus.Info("Get from DB")
			findAllResult = c.NewsRepository.FindAll(ctx, c.DB, request)

			// set to redis
			helper.SetRedis(c.Rds, key, findAllResult)
		}
	}

	return findAllResult
}
