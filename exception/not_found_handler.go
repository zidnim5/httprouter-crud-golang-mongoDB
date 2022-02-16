/***
* =======================================================================
* FILE NAME:        not_found_handler.go
* DATE CREATED:  	10-01-2022
* AUTHOR:			zidni.mujib5@gmail.com
* =======================================================================
 */

package exception

/**
 * NotFound struct
 */
type NotFound struct {
	error string
}

/**
 * Constructor
 */
func NewNotFoundError(err string) NotFound {
	return NotFound{error: err}
}
