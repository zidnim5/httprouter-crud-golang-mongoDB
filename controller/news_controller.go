/***
* =======================================================================
* FILE NAME:        news_controller.go
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
 * NewsController interface
 */
type NewsController interface {
	Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
