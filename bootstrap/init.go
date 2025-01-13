package bootstrap

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

func RegisterMetrics() {
	prometheus.MustRegister(UserSignups)
	log.Println("Prometheus metrics registered")
}

var (
	UserSignups = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_signups_total",
			Help: "Total number of user signups",
		},
		[]string{"status"}, // Labels: e.g., success or failure
	)
)
