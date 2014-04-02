package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	flagPort = flag.String("p", "12345", "port for http debug proxy")
)

func showUsage() {
	fmt.Fprintf(os.Stderr,
		"usage: %s [-p=<port>]\n",
		os.Args[0])
	fmt.Fprintf(os.Stderr,
		"flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = showUsage
	flag.Parse()

	http.HandleFunc("/", Proxy)
	fmt.Printf("HTTP Debug Proxy starting at %s...\n", *flagPort)
	err := http.ListenAndServe(":"+*flagPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
