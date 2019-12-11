package main

import (
	minimizers "github.com/adaptant-labs/go-minimizer"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minimization_service_requests_total",
		Help: "The total number of processed requests",
	})

	totalMinimizationRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minimization_service_minimization_requests_total",
		Help: "The total number of minimization requests",
	})

	totalAnonymizationRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "minimization_service_anonymization_requests_total",
		Help: "The total number of anonymization requests",
	})
)

func processRequestMetrics(level minimizers.MinimizationLevel) {
	totalRequests.Inc()

	if level == minimizers.MinimizationAnonymize {
		totalAnonymizationRequests.Inc()
	} else {
		totalMinimizationRequests.Inc()
	}
}
