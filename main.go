package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Proxy)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
