// gomocker project
package main

import (
	"flag"
	"fmt"
	"gomocker/utils"
	"log"
	"net/http"
)

var REQ_RESPS_PATH string = "reqResps/"
var mapping map[string]string

func main() {
	ip := flag.String("ip", "0.0.0.0", "ip on which to run mock service")
	port := flag.String("port", "9090", "port on which to run mock service")
	flag.Parse()

	mapping = utils.LoadReqResps(REQ_RESPS_PATH)

	router := http.NewServeMux()
	router.HandleFunc("/", logHandleRequest(handleGet))
	router.HandleFunc("/mvd1", logHandleRequest(handleMvd1))
	http.Handle("/", router)

	fmt.Println("GoMocks v1.0")
	fmt.Printf("Running on %s:%s", *ip, *port)

	err := http.ListenAndServe(*ip+":"+*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
