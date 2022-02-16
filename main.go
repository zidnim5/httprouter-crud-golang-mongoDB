/***
* =======================================================================
* FILE NAME:        main.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package main

import (
	"golang-mongodb/app"
	"golang-mongodb/controller"
	"golang-mongodb/exception"
	"golang-mongodb/helper"
	"golang-mongodb/repository"
	"golang-mongodb/services"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var (
	validate *validator.Validate
)

func main() {
	// todo: logrus for logs
	file, _ := os.OpenFile("storages/logs/application-"+time.Now().Format("02-01-2006")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logrus.SetOutput(file)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	//  todo: init
	DB := app.NewDB()
	Rds := app.NewRedis()

	validate := validator.New()
	tagsRepository := repository.NewTagsRepository()
	tagsService := services.NewTagsService(tagsRepository, DB, validate)
	tagsController := controller.NewTagsController(tagsService)
	newsRepository := repository.NewNewsRepository()
	newsService := services.NewNewsService(newsRepository, DB, Rds, validate)
	newsController := controller.NewNewsController(newsService)

	router := app.NewRouter(tagsController, newsController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	// notification console
	helper.ServiceReady()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
