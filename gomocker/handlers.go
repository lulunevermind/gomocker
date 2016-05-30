package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/jteeuwen/go-pkg-xmlx"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	resp := "SimpleMockServer v.1.0"
	fmt.Fprintf(w, resp)
}

func ByContainsString(w http.ResponseWriter, r *http.Request) {
	body := readBodyAsString(r)
	if strings.Contains(body, w.(DumpResponseWriter).tag) {
		resp := mapping[w.(DumpResponseWriter).template]
		fmt.Fprintf(w, resp)
	}
}

func GIBDD_simple_n_full(w http.ResponseWriter, r *http.Request) {
	body := readBodyAsString(r)
	if strings.Contains(body, "<dob:usrF>Èâàíîâ2</dob:usrF>") {
		resp := mapping["mvd_full.resp"]
		fmt.Fprintf(w, resp)
	}
	if strings.Contains(body, "<dob:usrF>simple</dob:usrF>") {
		resp := mapping["mvd_simple.resp"]
		fmt.Fprintf(w, resp)
	}
}

func FNS_LCBFindInfo_SendQINNNFL(w http.ResponseWriter, r *http.Request) {
	body := readBodyAsString(r)
	if strings.Contains(body, "<rq:inn>") {
		resp := mapping["fns_lcbfindinfo.resp"]
		fmt.Fprintf(w, resp)
	} else {
		resp := mapping["fns_qinn.resp"]
		fmt.Fprintf(w, resp)
	}
}

func ByXmlTagExists(w http.ResponseWriter, r *http.Request) {
	doc := xmlx.New()
	body := readBodyAsString(r)
	err := doc.LoadString(body, nil)
	if err != nil {
		panic(err)
	}
	node := doc.SelectNode("*", w.(DumpResponseWriter).tag)

	if node != nil {
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

func handleWithTemplateBy(fn http.HandlerFunc, tag string, template string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dumpRequestToLog(r)
		w.Header().Set("Content-type", "text/xml; charset=utf-8")
		dumpRespWriter := DumpResponseWriter{w, template, tag}
		Info.Println("TAG -->> ", tag)
		Info.Println("TEMPLATE -->> ", template)
		fn(dumpRespWriter, r)
	}
}

func handleWithCustomHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dumpRequestToLog(r)
		w.Header().Set("Content-type", "text/xml; charset=utf-8")
		dumpRespWriter := DumpResponseWriter{w, "", ""}
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
