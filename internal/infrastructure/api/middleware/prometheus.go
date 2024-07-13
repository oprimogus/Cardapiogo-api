package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

// PrometheusMetrics estrutura que armazena as métricas que queremos registrar.
type PrometheusMetrics struct {
	Registry          *prometheus.Registry
	RequestCounter    *prometheus.CounterVec
	ResponseTime      *prometheus.HistogramVec
	ErrorCounter      *prometheus.CounterVec
	ActiveConnections prometheus.Gauge
}

func NewPrometheusMetrics() *PrometheusMetrics {
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())

	requestCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cardapiogo_requests_total",
		Help: "Total de requisições recebidas.",
	}, []string{"path", "method", "status"})

	responseTime := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "cardapiogo_response_time_seconds",
		Help:    "Tempo de resposta da API.",
		Buckets: prometheus.DefBuckets, // Você pode personalizar os buckets
	}, []string{"path", "method", "status"})

	errorCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "cardapiogo_errors_total",
		Help: "Total de erros da API por endpoint, método e código de status.",
	}, []string{"path", "method", "status"})

	activeConnections := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cardapiogo_active_connections",
		Help: "Número atual de conexões ativas.",
	})

	reg.MustRegister(requestCounter, responseTime, errorCounter, activeConnections)

	return &PrometheusMetrics{
		Registry:          reg,
		RequestCounter:    requestCounter,
		ResponseTime:      responseTime,
		ErrorCounter:      errorCounter,
		ActiveConnections: activeConnections,
	}
}

func PrometheusMiddleware(metrics *PrometheusMetrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(v float64) {
			metrics.ResponseTime.WithLabelValues(path, method, fmt.Sprintf("%d", c.Writer.Status())).Observe(v)
		}))
		defer timer.ObserveDuration()

		metrics.ActiveConnections.Inc()
		defer metrics.ActiveConnections.Dec()

		// Processa o request
		c.Next()

		// Incrementa o contador de requests
		status := fmt.Sprintf("%d", c.Writer.Status())
		metrics.RequestCounter.WithLabelValues(path, method, status).Inc()

		// Incrementa o contador de erros, se necessário
		if c.Writer.Status() >= 400 {
			metrics.ErrorCounter.WithLabelValues(path, method, status).Inc()
		}
	}
}
