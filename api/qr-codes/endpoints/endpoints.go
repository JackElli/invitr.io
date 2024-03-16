package endpoints

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	qrcodes "invitr.io.com/qr-codes/endpoints/qr-codes"
	"invitr.io.com/responder"
)

type Endpoints struct {
	Logger *zap.Logger
}

func NewEndpointsMgr(logger *zap.Logger) *Endpoints {
	return &Endpoints{
		Logger: logger,
	}
}

func (e *Endpoints) SetupEndpoints(env string, r *mux.Router) error {
	responder := responder.NewResponder()

	public := r.PathPrefix("/").Subrouter()
	// add endpoints to the router
	_ = qrcodes.NewQRCodeMgr(public, e.Logger, responder)

	return nil
}
