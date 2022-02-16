/***
* =======================================================================
* FILE NAME:        news)controller_impl.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package controller

import (
	"golang-mongodb/helper"
	"golang-mongodb/model/web"
	"golang-mongodb/services"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type NewsControllerImpl struct {
	service services.NewsService
}

/**
 * Init
 */
func NewNewsController(services services.NewsService) NewsController {
	return &NewsControllerImpl{service: services}
}

func (c *NewsControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	newsCreateRequest := web.NewsCreateRequest{}
	helper.WriteToRequestBody(request, &newsCreateRequest)
	newsResponse := c.service.Create(request.Context(), newsCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Success: true,
		Status:  "Ok",
		Data:    newsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *NewsControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	newsUpdateRequest := web.NewsUpdateRequest{}
	helper.WriteToRequestBody(request, &newsUpdateRequest)

	newsResponse := c.service.Update(request.Context(), newsUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Success: true,
		Status:  "Ok",
		Data:    newsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *NewsControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	newsId := param.ByName("news_id")
	newsRequest := web.NewsDeleteRequest{
		Id: newsId,
	}

	c.service.Delete(request.Context(), newsRequest)

	webResponse := web.WebResponse{
		Code:    200,
		Success: true,
		Status:  "Data has been deleted",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *NewsControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	paramRequest := web.NewsFindRequest{
		Status: request.URL.Query().Get("status"),
		Topic:  request.URL.Query().Get("topic"),
	}

	newsResponse := c.service.FindAll(request.Context(), paramRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Success: true,
		Status:  "Ok",
		Data:    newsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
