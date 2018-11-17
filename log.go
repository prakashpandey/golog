/**
 * @author Prakash Pandey
 * @email prakashpandeyy@yahoo.com
 * @create date 2018-11-18 02:45:48
 * @modify date 2018-11-18 02:45:48
 * @desc [description]
 */
package main

import (
	"log"
)

const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
)

var LogLevel string

func Debug(message string, args ...interface{}) {
	if LogLevel == DEBUG {
		log.Fatalf(message, args)
	}
}
