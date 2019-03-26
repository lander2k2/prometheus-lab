package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "go_server_http_requests_total",
		Help: "Number of requests processed",
	})

	reqDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "go_server_response_duration",
		Help:    "Duration in seconds to process responses to requests",
		Buckets: []float64{10, 50, 100, 500},
	})
)

func main() {

	glog.Info("Starting go-server")

	http.Handle("/metrics", prometheus.Handler())
	glog.Infof("Prometheus metrics exposed at path /metrics")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	reqStart := time.Now()
	glog.Infof("Request received from %s", r.RemoteAddr)

	// increment requests count for prometheus metrics
	httpRequests.Inc()

	respContent := `
{
  "data": {
    "message": "default"
  }
}
`
	fmt.Fprintf(w, respContent)
	glog.Infof("Returned default response to %s", r.RemoteAddr)
	duration := time.Now().Sub(reqStart)

	// measure time taken to process request
	reqDuration.Observe(float64(duration / time.Millisecond))
}
