/***
* =======================================================================
* FILE NAME:        error_handler.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package exception

import (
	"golang-mongodb/helper"
	"golang-mongodb/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

/**
 * todo: error handler
 */
func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	if writeError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

/**
 * todo: validationError handler
 */
func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {

	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Success: false,
			Status:  "BAD REQUEST",
			Data:    exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

/**
 * todo: notFoundError handler
 */
func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(NotFound)

	if ok {
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Success: false,
			Status:  "NOT FOUND",
			Data:    err,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

/**
 * todo: writerError handler
 */
func writeError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(mongo.WriteException)
	if ok {

		writer.WriteHeader(http.StatusConflict)

		webResponse := web.WebResponse{
			Code:    409,
			Success: false,
			Status:  "DATA IS EXIST",
			Data:    err,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

/**
 * todo: internalServerError handler
 */
func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	logrus.Error(err)
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
