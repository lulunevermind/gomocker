package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kylewolfe/simplexml"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		resp := "SimpleMockServer v.1.0"
		fmt.Fprintf(w, resp) // send data to client side
	}
}

func handleMvd1(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		body := readBodyAsString(r)
		if strings.Contains(body, "<deptcode>") {
			resp := mapping["mvd1.resp"]
			fmt.Fprintf(w, resp)
		}
	}
}

func logHandleRequest(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form)
		//		fmt.Println(readBodyAsString(r))
		fmt.Println(r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
		fn(w, r)
	}
}

func readBodyAsXml(req *http.Request) *simplexml.Document {
	bd, _ := ioutil.ReadAll(req.Body)
	as_str := string(bd)
	xml_doc, err := simplexml.NewDocumentFromReader(strings.NewReader(as_str))
	if err != nil {
		panic(err)
	}
	return xml_doc
}

func readBodyAsString(req *http.Request) string {
	bd, _ := ioutil.ReadAll(req.Body)
	return string(bd)
}
