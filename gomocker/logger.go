package main

import "log"
import "io"

var (
	Trace *log.Logger
	Info  *log.Logger
)

func Init_logger(
	traceHandle io.Writer,
	infoHandle io.Writer) {

	Trace = log.New(traceHandle,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
