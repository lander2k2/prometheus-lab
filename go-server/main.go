package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

func main() {

	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.Info("Starting go-server")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("Request received from %s", r.RemoteAddr)

	respContent := `
{
  "data": {
    "message": "default"
  }
}
`
	fmt.Fprintf(w, respContent)
	glog.Infof("Returned default response to %s", r.RemoteAddr)
}
