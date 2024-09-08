package main

import (
    "crypto/rand"
    "flag"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    maxSize     int
    srvPort     int
    keyLengthHist = prometheus.NewHistogram(prometheus.HistogramOpts{
        Name:    "key_length_distribution",
        Help:    "Key length distribution",
        Buckets: prometheus.LinearBuckets(0, float64(maxSize)/20, 20),
    })
    statusCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_status_codes",
            Help: "HTTP status codes",
        },
        []string{"code"},
    )
)

func init() {
    // Register Prometheus metrics
    prometheus.MustRegister(keyLengthHist)
    prometheus.MustRegister(statusCounter)
}

func keyHandler(w http.ResponseWriter, r *http.Request) {
    keyLengthStr := r.URL.Query().Get("length")
    if keyLengthStr == "" {
        http.Error(w, "Missing length parameter", http.StatusBadRequest)
        statusCounter.WithLabelValues("400").Inc()
        return
    }

    keyLength, err := strconv.Atoi(keyLengthStr)
    if err != nil || keyLength > maxSize {
        http.Error(w, "Invalid key length or exceeds max size", http.StatusBadRequest)
        statusCounter.WithLabelValues("400").Inc()
        return
    }

    key := make([]byte, keyLength)
    _, err = rand.Read(key)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        statusCounter.WithLabelValues("500").Inc()
        return
    }

    w.Write(key)
    keyLengthHist.Observe(float64(keyLength))
    statusCounter.WithLabelValues("200").Inc()
}

func main() {
    // Command-line arguments
    flag.IntVar(&maxSize, "max-size", 1024, "maximum key size")
    flag.IntVar(&srvPort, "srv-port", 1123, "server listening port")
    flag.Parse()

    http.HandleFunc("/key", keyHandler)
    http.Handle("/metrics", promhttp.Handler())

    log.Printf("Starting server on port %d with max key size %d...", srvPort, maxSize)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", srvPort), nil))
}