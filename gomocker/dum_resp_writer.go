package main

import (
	"net/http"
)

// First, we wrap the ResponseWriter interface with our own response writer, that will log everything and pass
// all functionality to a real response writer:

type DumpResponseWriter struct {
	w        http.ResponseWriter
	template string
	tag      string
}

func (w DumpResponseWriter) Header() http.Header {
	return w.w.Header()
}

func (w DumpResponseWriter) Write(b []byte) (int, error) {
	// You can add more context about the connection when initializing the writer and log it here
	Info.Println("Response -->> ", w)
	Info.Println(string(b))
	return w.w.Write(b)
}

func (w DumpResponseWriter) WriteHeader(h int) {
	Info.Println("Response HEADER -->> ", h)
	Info.Println(string(h))
	w.w.WriteHeader(h)
}

// The handler function is the same and agnostic to the fact that we're using a  "Fake" writer...

func MyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
