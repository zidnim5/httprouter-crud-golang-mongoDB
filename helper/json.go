/***
* =======================================================================
* FILE NAME:        json.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package helper

import (
	"encoding/json"
	"golang-mongodb/model/web"
	"net/http"
)

/**
 * todo: for convert requestBody to struct parameter
 */
func WriteToRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&result)
	PanicIfError(err)
}

/**
 * todo: toggle above func
 */
func WriteToResponseBody(writer http.ResponseWriter, webResponse web.WebResponse) {
	writer.Header().Add("Content-Type", "Application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}
