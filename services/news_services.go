/***
* =======================================================================
* FILE NAME:        news_service.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package services

import (
	"context"
	"golang-mongodb/model/web"
)

/**
 * NewsService interface
 */
type NewsService interface {
	Create(ctx context.Context, request web.NewsCreateRequest) web.NewsResponse
	Update(ctx context.Context, request web.NewsUpdateRequest) web.NewsResponse
	Delete(ctx context.Context, request web.NewsDeleteRequest) interface{}
	FindAll(ctx context.Context, request web.NewsFindRequest) interface{}
}
