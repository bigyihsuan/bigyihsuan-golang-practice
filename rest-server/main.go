package main

import (
	"net/http"
	"os"
	"rest-server/router"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
	router := router.ConfigureRouter()
	addr := ":3000"

	logrus.WithField("addr", addr).Info("Starting server...")

	if err := http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, router)); err != nil {
		logrus.WithField("event", "start server").Fatal(err)
	}
}
