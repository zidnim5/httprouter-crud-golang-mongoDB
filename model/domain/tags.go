/***
* =======================================================================
* FILE NAME:        tags.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package domain

import "encoding/json"

/**
 * Tags struct
 */
type Tags struct {
	Id   string `json:"_id" bson:"_id"`
	Name string `json:"name"`
}

func (t Tags) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(t)
	return bytes, err
}
