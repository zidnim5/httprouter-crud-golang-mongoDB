/***
* =======================================================================
* FILE NAME:        tags_controller_impl.go
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

type TagsControllerImpl struct {
	service services.TagsService
}

/**
 * Init
 */
func NewTagsController(services services.TagsService) TagsController {
	return &TagsControllerImpl{service: services}
}

func (c *TagsControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	tagsCreateRequest := web.TagsRequest{}
	helper.WriteToRequestBody(request, &tagsCreateRequest)
	tagsResponse := c.service.Create(request.Context(), tagsCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Success: true,
		Status:  "Ok",
		Data:    tagsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *TagsControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {

	tagsUpdateRequest := web.TagsUpdateRequest{}
	helper.WriteToRequestBody(request, &tagsUpdateRequest)

	tagsResponse := c.service.Update(request.Context(), tagsUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Success: true,
		Status:  "Ok",
		Data:    tagsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *TagsControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	tagsId := param.ByName("tags_id")
	tagsRequest := web.TagsDeleteRequest{
		Id: tagsId,
	}

	c.service.Delete(request.Context(), tagsRequest)

	webResponse := web.WebResponse{
		Code:    200,
		Success: true,
		Status:  "Ok",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (c *TagsControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	tagsResponse := c.service.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:    200,
		Success: true,
		Status:  "Ok",
		Data:    tagsResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
