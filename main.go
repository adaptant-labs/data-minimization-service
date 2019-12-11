package main

import (
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	r := NewServiceRouter()
	l := handlers.LoggingHandler(os.Stdout, r)

	log.Info("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", l))
}
