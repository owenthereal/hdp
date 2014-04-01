package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ProxyServer(w http.ResponseWriter, req *http.Request) {
	req.Header.Set("Host", req.Host)

	fmt.Printf("> %s %s %s\n", req.Method, reqPath(req), req.Proto)

	for k, v := range req.Header {
		fmt.Printf("> %s - %s\n", k, strings.Join(v, ","))
	}

	content, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err == nil {
		fmt.Println(string(content))
	}

	fmt.Println("")

	fmt.Fprintln(w, "ok")
}

func reqPath(req *http.Request) string {
	path := req.URL.Path
	query := []string{}
	for k, v := range req.URL.Query() {
		query = append(query, fmt.Sprintf("%s=%s", k, strings.Join(v, ",")))
	}

	if len(query) > 0 {
		path = fmt.Sprintf("%s?%s", path, strings.Join(query, "&"))
	}

	return path
}

func main() {
	http.HandleFunc("/", ProxyServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
