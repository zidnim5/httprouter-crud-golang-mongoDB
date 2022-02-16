/***
* =======================================================================
* FILE NAME:        tags_service_impl.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package services

import (
	"context"
	"golang-mongodb/exception"
	"golang-mongodb/helper"
	"golang-mongodb/model/domain"
	"golang-mongodb/model/web"
	"golang-mongodb/repository"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	DB             *mongo.Client
	validate       *validator.Validate
}

/**
 * Init
 */
func NewTagsService(tagsRepository repository.TagsRepository, DB *mongo.Client, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{TagsRepository: tagsRepository, DB: DB, validate: validate}
}

/**
 * todo: create new tags
 *
 * @param context.Context
 * @param web.TagsRequest
 *
 * @return web.TagsResponse
 */
func (c *TagsServiceImpl) Create(ctx context.Context, request web.TagsRequest) web.TagsResponse {
	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	created := c.TagsRepository.Save(ctx, c.DB, request)

	result := web.TagsResponse{
		Id:   created.Id,
		Name: created.Name,
	}

	return result

}

/**
 * todo: update tags
 *
 * @param context.Context
 * @param web.TagsUpdateRequest
 *
 * @return web.TagsResponse
 */
func (c *TagsServiceImpl) Update(ctx context.Context, request web.TagsUpdateRequest) web.TagsResponse {
	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	objectIDS, _ := primitive.ObjectIDFromHex(request.Id)
	filter := bson.M{"_id": objectIDS}
	changes := web.TagsRequest{
		Name: request.Name,
	}

	updatedResult := c.TagsRepository.Update(ctx, c.DB, filter, changes)

	if updatedResult {
		return web.TagsResponse{
			Id:   request.Id,
			Name: request.Name,
		}
	}
	panic(exception.NewNotFoundError("Data not found"))
}

/**
 * todo: delete tags
 *
 * @param context.Context
 * @param web.TagsDeleteRequest
 *
 */
func (c *TagsServiceImpl) Delete(ctx context.Context, request web.TagsDeleteRequest) interface{} {
	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	objectIDS, _ := primitive.ObjectIDFromHex(request.Id)
	selector := bson.M{"_id": objectIDS}

	deletedResult := c.TagsRepository.Delete(ctx, c.DB, selector)
	if deletedResult {
		return nil
	}
	panic(exception.NewNotFoundError("Data not found"))
}

/**
 * todo: delete tags
 *
 * @param context.Context
 *
 * @return []domain.Tags
 *
 */
func (c *TagsServiceImpl) FindAll(ctx context.Context) []domain.Tags {
	findAllResult := c.TagsRepository.FindAll(ctx, c.DB)

	return findAllResult
}
