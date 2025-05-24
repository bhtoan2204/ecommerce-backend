package http

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type AppMetric interface {
	Handler() http.Handler
}

type appMetric struct {
}

func NewMetric() (AppMetric, error) {
	return &appMetric{}, nil
}

func (m *appMetric) Handler() http.Handler {
	return promhttp.Handler()
}
