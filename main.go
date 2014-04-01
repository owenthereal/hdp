package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ProxyServer(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Printf("%s - %s\n", k, strings.Join(v, ","))
	}

	fmt.Fprintln(w, "ok")
}

func main() {
	http.HandleFunc("/", ProxyServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
