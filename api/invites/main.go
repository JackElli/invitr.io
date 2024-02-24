package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitio.com/invites/endpoints"
)

// Need to make sure that database is set up correctly

// This microservice holds all of the info about invites
// it should be able to create, delete, edit invites
// add QR codes and retrieve info from the users microservice.

// The idea is that a user creates an invite to an event,
// which either:
// 1. they send to the recepient via email (either with QR code or direct to
//    calender)
// 2. they print out invites on a special card (for weddings, etc...)

// then the recipient can say whether they are going, add to calender,
// add a message etc..

// The event should hold info like date, place, who is organising etc..
// We could also add links to buy things like flowers, drinks, food
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

	logger.Info("Started invites api...")
	err = http.ListenAndServe(":3202", http.Handler(r))
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
