package diagnostics

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func NewServer(log *logrus.Logger, diagPort string) *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz(log)).Methods(http.MethodGet)
	router.HandleFunc("/readyz", Readyz(log)).Methods(http.MethodGet)

	diagServer := &http.Server{
		Addr:    ":" + diagPort,
		Handler: router,
	}

	return diagServer
}
