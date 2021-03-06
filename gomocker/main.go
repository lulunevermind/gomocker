// gomocker project
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var REQ_RESPS_PATH string = "reqResps/"
var mapping map[string]string

func main() {
	port := flag.String("port", "8081", "port on which to run mock service")
	//	logs := flag.String("log", "nolog", "stdout or no logs")
	logs := flag.Bool("debug", false, "show debug info in stdout")
	flag.Parse()

	if *logs == false {
		Init_logger(ioutil.Discard, ioutil.Discard)
	} else {
		Init_logger(os.Stdout, os.Stdout)
	}

	mapping = LoadReqResps(REQ_RESPS_PATH)

	router := mux.NewRouter()
	router.HandleFunc("/", logHandleRequestStrictIn(handleGet, 1)).Methods("GET")
	router.HandleFunc("/mvd1", handleWithTemplateBy(ByXmlTagExists, "<deptcode>", "mvd1.resp")).Methods("POST")
	router.HandleFunc("/SID0003030", handleWithCustomHandler(GIBDD_simple_n_full)).Methods("POST")
	router.HandleFunc("/SID0003957", handleWithTemplateBy(ByContainsString, "<smev:Code>FSSP01001</smev:Code>", "fssp.resp")).Methods("POST")
	router.HandleFunc("/SID0003049", handleWithCustomHandler(FNS_LCBFindInfo_SendQINNNFL)).Methods("POST")

	fmt.Println("GoMocks v1.0")
	fmt.Printf("Running on %s:%s", "0.0.0.0", *port)
	fmt.Println("")

	err := http.ListenAndServe(":"+*port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
