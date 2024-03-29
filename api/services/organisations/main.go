package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitr.io.com/cors"
	"invitr.io.com/services/organisations/endpoints"
)

// This microservice should hold information about
// each organisation. I want invitr.io to be able to
// invite people within your organisation to meetings
// very easily. We already have the backbone for this
// in place with invite ids and keys, we just need a way
// to switch it so we use the user id instead of username
// then we can invite people through invitr.io, not just
// on an email.
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

	logger.Info("Started organisations api...")
	err = http.ListenAndServe(":3203", cors.CORS(r, ENVIRONMENT))
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
