/***
* =======================================================================
* FILE NAME:        redis.go
* DATE CREATED:  	11-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */
package app

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println("Connected to Redis")
	fmt.Println("=======================")

	return client
}
