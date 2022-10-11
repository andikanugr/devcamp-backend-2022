# Tokopedia Devcamp 2022

This repository holds the code examples that will be used throughout the Tokopedia Devcamp 2022 instructor-led sessions.

## Running the Code

To run the whole project, you can use

```bash
docker-compose up
```

### Export metric
```go
// Prometheus endpoint
router.Path("/prometheus").Handler(promhttp.Handler())
```

### Add prometheus metrics
```go
var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"path", "status"},
)

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_response_time_seconds",
	Help: "Duration of HTTP requests.",
}, []string{"path"})


prometheus.Register(totalRequests)
prometheus.Register(responseStatus)
prometheus.Register(httpDuration)

```

### Add metrics middleware
```go
// Middleware to collecting the metrics of the http request
func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode

		responseStatus.WithLabelValues(path, strconv.Itoa(statusCode)).Inc()
		totalRequests.WithLabelValues(path).Inc()

		timer.ObserveDuration()
	})
}

router.Use(prometheusMiddleware)
```

### Let's build dashboard
Open http://localhost:3000 to access grafana dashboard.

### Load Test
Let's check that we can monitor our API.
```bash
echo "GET http://localhost:9000/" | vegeta attack -duration=5s | vegeta report
echo "GET http://localhost:9000/user/1" | vegeta attack -duration=5s | vegeta report
echo "GET http://localhost:9000/user/2" | vegeta attack -duration=5s | vegeta report
```