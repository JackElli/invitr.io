package qrcodes

import (
	"qr-codes/responder"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ROOT = "/qr-code"
)

type QRCodeMgr struct {
	Logger    *zap.Logger
	Router    *mux.Router
	Responder responder.Responder
}

func NewQRCodeMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder) *QRCodeMgr {
	e := &QRCodeMgr{
		Logger:    logger,
		Router:    router,
		Responder: responder,
	}
	e.Register()
	return e
}

func (mgr *QRCodeMgr) Register() {

}
