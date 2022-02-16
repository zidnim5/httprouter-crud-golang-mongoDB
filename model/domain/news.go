/***
* =======================================================================
* FILE NAME:        news.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package domain

/**
 * News struct
 */
type News struct {
	Id      string   `json:"_id" bson:"_id"`
	Topic   string   `json:"topic" bson:"topic"`
	Title   string   `json:"title" bson:"title"`
	Status  string   `json:"status" bson:"status"`
	Tags    []string `json:"tags" bson:"tags"`
	Content string   `json:"content" bson:"content"`
}
