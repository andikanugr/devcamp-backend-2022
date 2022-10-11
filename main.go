package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

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

func init() {
	prometheus.Register(totalRequests)
	prometheus.Register(responseStatus)
	prometheus.Register(httpDuration)

	// log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.DebugLevel)
}

func main() {
	router := mux.NewRouter()
	router.Use(prometheusMiddleware)

	// Prometheus endpoint
	router.Path("/prometheus").Handler(promhttp.Handler())

	// Serving static files
	router.HandleFunc("/user/{id:[0-9]+}", handleGetUser).Methods(http.MethodGet)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./files/")))

	fmt.Println("Serving requests on port 9000")
	log.WithFields(log.Fields{
		"port": 9000,
	}).Info("Start listener")
	err := http.ListenAndServe(":9000", router)
	log.Fatal(err)
}

type User struct {
	ID    int64
	Name  string
	Email string
}

func handleGetUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		log.WithFields(log.Fields{"vars": vars["id"]}).Error("Failed to get param")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failed to encode response")
		return
	}

	user, err := getUserData(id)
	log.WithFields(log.Fields{"user": user}).Debug("getUserData")
	if err != nil {
		log.WithFields(log.Fields{"userId": id}).Error(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err.Error())
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		log.Error("Failed to encode response")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to encode response")
		return
	}
	fmt.Fprint(w, string(response))
}

func getUserData(userId int64) (*User, error) {
	if userId == 2 {
		return nil, errors.New("user not found")
	}

	dummyUser := &User{
		ID:    userId,
		Name:  "John Doe",
		Email: "john.doe@gmail.com",
	}
	return dummyUser, nil
}
