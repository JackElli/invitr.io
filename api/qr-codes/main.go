package main

import (
	"log"
	"net/http"
	"qr-codes/endpoints"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// This microservice doesn't need to persist anything
// and should only create and return QR codes
// that should be stored in the invites microservice
func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	const ENVIRONMENT = "dev"

	r := mux.NewRouter()

	endpoints := endpoints.NewEndpointsMgr(logger)
	err := endpoints.SetupEndpoints(ENVIRONMENT, r)
	if err != nil {
		logger.Error("Cannot setup endpoints", zap.Error(err))
		return
	}

	logger.Info("Started qr-codes api...")
	err = http.ListenAndServe(":3201", http.Handler(r))
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
