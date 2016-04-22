// gomocker project
package main

import (
	"flag"
	"fmt"
	"gomocker/gomocker/utils"
	"log"
	"net/http"
	"os"
)

var REQ_RESPS_PATH string = "reqResps/"
var mapping map[string]string

func main() {
	ip := flag.String("ip", "0.0.0.0", "ip on which to run mock service")
	port := flag.String("port", "9090", "port on which to run mock service")
	flag.Parse()

	mapping = utils.LoadReqResps(REQ_RESPS_PATH)

	router := http.NewServeMux()
	router.HandleFunc("/", logHandleRequestStrictIn(handleGet, 1))
	router.HandleFunc("/mvd1", logHandleRequestInDelta(handleMvd1, 1, 5))
	http.Handle("/", router)

	fmt.Println("GoMocks v1.0")
	fmt.Printf("Running on %s:%s", *ip, *port)
	fmt.Println("")
	Init_logger(os.Stdout, os.Stdout, os.Stdout, os.Stdout)

	err := http.ListenAndServe(*ip+":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
