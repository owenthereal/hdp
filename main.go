package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	flagHost = flag.String("h", "localhost", "host")
	flagPort = flag.String("p", "12345", "port")
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
	fmt.Println(Figlet)
	fmt.Printf("HDP listening on %s:%s\n\n", *flagHost, *flagPort)
	err := http.ListenAndServe(*flagHost+":"+*flagPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
