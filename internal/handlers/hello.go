package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Hello(log *logrus.Logger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		log.Infof("%v", values)
		//time.Sleep(5 * time.Second)
		w.WriteHeader(http.StatusOK)
	}
}
