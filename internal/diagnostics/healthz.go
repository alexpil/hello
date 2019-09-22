package diagnostics

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Healthz(log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		log.Info("Healthz was called")
		w.WriteHeader(http.StatusOK)
	}
}

func Readyz(log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		log.Info("Healthz was called")
		w.WriteHeader(http.StatusOK)
	}
}
