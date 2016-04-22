package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/kylewolfe/simplexml"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		resp := "SimpleMockServer v.1.0"
		fmt.Fprintf(w, resp)
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

func logHandleRequestStrictIn(fn http.HandlerFunc, seconds time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(seconds * time.Second)
		dumpRequestToLog(r)
		fn(w, r)
		Info.Printf("Respond in %s", seconds)
		Info.Println()
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func logHandleRequestInDelta(fn http.HandlerFunc, d_min time.Duration, d_max time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		delta := random(int(d_min), int(d_max))
		time.Sleep(time.Duration(delta) * time.Second)
		dumpRequestToLog(r)
		fn(w, r)
		Info.Printf("Respond in %s", string(delta))
		Info.Println()
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

func dumpRequestToLog(r *http.Request) {
	dumped, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}
	Info.Println(string(dumped))
}
