/***
* =======================================================================
* FILE NAME:        database.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package app

import (
	"context"
	"fmt"
	"golang-mongodb/helper"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/**
 * todo: get connection to mongo db
 *
 * @return *mongo.Client
 */
func NewDB() *mongo.Client {
	fmt.Println("=======================")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	_ = cancel

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	helper.PanicIfError(err)
	// connect to mongo
	err = client.Connect(ctx)
	helper.PanicIfError(err)
	// test connection to mongo
	err = client.Ping(ctx, readpref.Primary())
	helper.PanicIfError(err)
	fmt.Println("Connected to mongoDB")

	return client
}
