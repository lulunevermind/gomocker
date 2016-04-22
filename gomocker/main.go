// gomocker project
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var REQ_RESPS_PATH string = "reqResps/"
var mapping map[string]string

func main() {
	port := flag.String("port", "9090", "port on which to run mock service")
	logs := flag.String("log", "nolog", "stdout or no logs")
	flag.Parse()

	mapping = LoadReqResps(REQ_RESPS_PATH)

	if *logs == "nolog" {
		Init_logger(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)
	} else {
		Init_logger(os.Stdout, os.Stdout, os.Stdout, os.Stdout)
	}

	router := http.NewServeMux()
	router.HandleFunc("/", logHandleRequestStrictIn(handleGet, 1))
	router.HandleFunc("/mvd1", logHandleRequest(handleMvd1))
	router.HandleFunc("/mvd2", logHandleRequest(handleMvd2))
	http.Handle("/", router)

	fmt.Println("GoMocks v1.0")
	fmt.Printf("Running on %s:%s", "0.0.0.0", *port)
	fmt.Println("")

	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
