package responder

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type Responder interface {
	Respond(w http.ResponseWriter, code int, response interface{})
	Error(w http.ResponseWriter, code int, err error)
}

type Respond struct{}

// NewResponder returns a respond struct implementing the
// responder interface
func NewResponder() *Respond {
	return &Respond{}
}

// Send a JSON response back
func (r *Respond) Respond(w http.ResponseWriter, code int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	// if theres no response, send an empty object
	if response == nil || reflect.ValueOf(response).IsNil() {
		json.NewEncoder(w).Encode(make(map[string]interface{}))
		return
	}

	json.NewEncoder(w).Encode(response)
}

// Send an error JSON response back
func (r *Respond) Error(w http.ResponseWriter, code int, err error) {
	r.Respond(w, code, map[string]interface{}{
		"Error": err.Error(),
	})
}
