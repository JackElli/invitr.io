package cors

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// GetAllowedMethods returns the CORS methods we allow
func GetAllowedMethods() []string {
	return []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
}

// GetAllowedHeaders returns the CORS headers we allow
func GetAllowedHeaders() []string {
	return []string{"Origin", "Content-Length", "Content-Type"}
}

// GetAllowedOrigin returns the CORS origins we allow the backend
// to receive from
func GetAllowedOrigin(env string) []string {
	// TODO make these env variables
	if env != "prod" {
		return []string{"http://localhost:3000"}
	}

	// return the prod origin if on prod
	return []string{""}
}

// GetExposedHeaders returns the headers we want to expose
func GetExposedHeaders() []string {
	return []string{"Set-Cookie"}
}

// setCorsMethods allows us to choose which headers are allowed
func CORS(r *mux.Router, env string) http.Handler {
	return handlers.CORS(
		handlers.AllowedMethods(
			GetAllowedMethods(),
		),
		handlers.AllowedOrigins(
			GetAllowedOrigin(env),
		),
		handlers.ExposedHeaders(
			GetExposedHeaders(),
		),
		handlers.AllowedHeaders(
			GetAllowedHeaders(),
		),
		handlers.AllowCredentials(),
	)(r)
}
