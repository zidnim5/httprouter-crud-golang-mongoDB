/***
* =======================================================================
* FILE NAME:        tags_controller.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/**
 * TagsController interface
 */
type TagsController interface {
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
