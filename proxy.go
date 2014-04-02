package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Proxy(w http.ResponseWriter, req *http.Request) {
	// put back Host to the header
	req.Header.Set("Host", req.Host)

	displayReq(req)
	displayHeader(req)
	displayBody(req)
	fmt.Println("")

	fmt.Fprintln(w, "ok")
}

func displayReq(req *http.Request) {
	display("%s %s %s\n", req.Method, reqPath(req), req.Proto)
}

func displayHeader(req *http.Request) {
	for k, v := range req.Header {
		display("%s: %s\n", k, strings.Join(v, ","))
	}
}

func displayBody(req *http.Request) {
	content, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err == nil {
		fmt.Println(string(content))
	}
}

func display(format string, a ...interface{}) {
	format = fmt.Sprintf("> %s", format)
	fmt.Printf(format, a...)
}

func reqPath(req *http.Request) string {
	path := req.URL.Path
	query := []string{}
	for k, v := range req.URL.Query() {
		for _, vv := range v {
			query = append(query, fmt.Sprintf("%s=%s", k, vv))
		}
	}

	if len(query) > 0 {
		path = fmt.Sprintf("%s?%s", path, strings.Join(query, "&"))
	}

	return path
}
