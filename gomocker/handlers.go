package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	resp := "SimpleMockServer v.1.0"
	fmt.Fprintf(w, resp)
}

func ByContainsTag(w http.ResponseWriter, r *http.Request) {
	body := readBodyAsString(r)
	if strings.Contains(body, w.(DumpResponseWriter).tag) {
		resp := mapping[w.(DumpResponseWriter).template]
		fmt.Fprintf(w, resp)
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
		Info.Printf("Respond in %s", delta)
		Info.Println()
	}
}

func handleByTagWithTemplateLogged(fn http.HandlerFunc, tag string, template string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dumpRequestToLog(r)
		dumpRespWriter := DumpResponseWriter{w, template, tag}
		Info.Println("TAG -->> ", tag)
		Info.Println("TEMPLATE -->> ", template)
		fn(dumpRespWriter, r)
	}
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
