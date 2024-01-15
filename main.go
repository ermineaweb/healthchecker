package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var path string
var port int

func init() {
	flag.IntVar(&port, "port", 3000, "the target application port")
	flag.StringVar(&path, "path", "/healthcheck", "the path for the healthcheck")
	flag.Parse()
}
func main() {
	if res, err := http.Get(fmt.Sprintf("http://localhost:%v%v", port, path)); err != nil || res.StatusCode > 299 {
		os.Exit(1)
	}
}
