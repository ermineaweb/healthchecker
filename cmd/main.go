package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Headers []string

var (
	port    int
	path    string
	headers Headers
)

func init() {
	flag.IntVar(&port, "port", 3000, "target application port")
	flag.StringVar(&path, "path", "/healthcheck", "path for the healthcheck")
	flag.Var(&headers, "header", "request headers")
	flag.Parse()
}

func main() {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%v%v", port, path), nil)
	if err != nil {
		os.Exit(1)
	}

	for _, header := range headers {
		fmt.Println("set header")
		req.Header.Set("x-token", header)
	}

	client := &http.Client{}
	if res, err := client.Do(req); err != nil || res.StatusCode > 299 {
		os.Exit(1)
	} else {
		fmt.Println("ok")
	}
}

func (he *Headers) String() string {
	return "my string representation"
}

func (he *Headers) Set(value string) error {
	*he = append(*he, value)
	return nil
}
