package main

import (
	"fmt"
	//	"io/ioutil"
	"net/http"
	"strings"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		resp := "SimpleMockServer v.1.0"
		fmt.Fprintf(w, resp) // send data to client side
	}
}

func handleMvd1(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		resp := mapping["mvd1.resp"]
		fmt.Fprintf(w, resp)
	}
}

func logHandleRequest(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()       // parse arguments, you have to call this by yourself
		fmt.Println(r.Form) // print form information in server side
		fmt.Println(r.Method, r.URL.Path)
		//		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
		fn(w, r)
	}
}
