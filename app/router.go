/***
* =======================================================================
* FILE NAME:        redis.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package app

import (
	"golang-mongodb/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(tagsController controller.TagsController, newsController controller.NewsController) *httprouter.Router {
	router := httprouter.New()

	// tags
	router.GET("/api/tags", tagsController.FindAll)
	router.POST("/api/tags", tagsController.Create)
	router.PATCH("/api/tags", tagsController.Update)
	router.DELETE("/api/tags/:tags_id", tagsController.Delete)

	// news
	router.GET("/api/news", newsController.FindAll)
	router.POST("/api/news", newsController.Create)
	router.PATCH("/api/news", newsController.Update)
	router.DELETE("/api/news/:news_id", newsController.Delete)

	return router

}
