package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitr.io.com/services/users/endpoints"
)

// This microservice holds all of the info about the users
// currently in the system

// it should be able to add, remove, edit users
// and should be able to pass info across different
// microservices when required.
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

	logger.Info("Started users api...")
	err = http.ListenAndServe(":3200", http.Handler(r))
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
