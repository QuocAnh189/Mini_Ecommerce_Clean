package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddleware() gin.HandlerFunc {
	requests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests processed",
		},
		[]string{"method", "path"},
	)

	latency := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response time",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	prometheus.MustRegister(requests, latency)

	return func(c *gin.Context) {
		timer := prometheus.NewTimer(latency.WithLabelValues(c.Request.Method, c.Request.URL.Path))
		c.Next()
		timer.ObserveDuration()

		requests.WithLabelValues(c.Request.Method, c.Request.URL.Path).Inc()
	}
}
