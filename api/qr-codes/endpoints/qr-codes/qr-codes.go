package qrcodes

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitio.com/qr-codes/responder"
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

func (mgr *QRCodeMgr) Generate() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		qrcode := QRCode{
			Width:  200,
			Height: 200,
			Pixels: []int{1, 2, 3},
		}

		mgr.Logger.Info("QR code generated successfully.")
		mgr.Responder.Respond(w, 201, qrcode)
	}
}

func (mgr *QRCodeMgr) Register() {
	mgr.Router.HandleFunc(ROOT, mgr.Generate()).Methods("GET")
}
