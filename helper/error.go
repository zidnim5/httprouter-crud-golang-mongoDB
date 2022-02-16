/***
* =======================================================================
* FILE NAME:        error.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package helper

import "github.com/sirupsen/logrus"

/**
 * todo: for panic
 */
func PanicIfError(err error) {
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
}
