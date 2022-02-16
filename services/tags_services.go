/***
* =======================================================================
* FILE NAME:        tags_service.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package services

import (
	"context"
	"golang-mongodb/model/domain"
	"golang-mongodb/model/web"
)

/**
 * TagsService interface
 */
type TagsService interface {
	Create(ctx context.Context, request web.TagsRequest) web.TagsResponse
	Update(ctx context.Context, request web.TagsUpdateRequest) web.TagsResponse
	Delete(ctx context.Context, request web.TagsDeleteRequest) interface{}
	FindAll(ctx context.Context) []domain.Tags
}
