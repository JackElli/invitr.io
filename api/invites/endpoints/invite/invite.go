package invite

import (
	"encoding/json"
	"invites/invitesmgr"
	"invites/responder"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	QR_CODE_URL = "http://qr-codes:3201/qr-code"
	ROOT        = "/invite"
	INVITE      = ROOT + "/{inviteId}"
)

type InviteMgr struct {
	Logger      *zap.Logger
	Router      *mux.Router
	Responder   responder.Responder
	InviteStore invitesmgr.InviteStorer
}

func NewInviteMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, invitestore invitesmgr.InviteStorer) *InviteMgr {
	e := &InviteMgr{
		Logger:      logger,
		Router:      router,
		Responder:   responder,
		InviteStore: invitestore,
	}
	e.Register()
	return e
}

// generateQRCode fetches a QR code from the QR code microservice
func generateQRCode() (*invitesmgr.QRCode, error) {
	resp, err := http.Get(QR_CODE_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var qrCode invitesmgr.QRCode
	json.NewDecoder(resp.Body).Decode(&qrCode)

	return &qrCode, nil
}

// NewInvite creates a new invite based on some user input
func (mgr *InviteMgr) NewInvite() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var getinvite invitesmgr.Invite
		json.NewDecoder(req.Body).Decode(&getinvite)

		// we need to generate a QR code, for this we need to
		// call the QR code microservice
		qrcode, err := generateQRCode()
		if err != nil {
			mgr.Responder.Error(w, 500, err)
			return
		}

		// We could possibly send it as a byte array
		// but we'll need to unmarshal it properly
		qrcodeBytes, _ := json.Marshal(*qrcode)
		getinvite.QRCode = string(qrcodeBytes)
		invite, err := mgr.InviteStore.Insert(&getinvite)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		inviteJSON := invitesmgr.InviteJSON{
			Id:         invite.Id,
			Organiser:  invite.Organiser,
			Location:   invite.Location,
			Date:       invite.Date,
			QRCode:     *qrcode,
			Passphrase: invite.Passphrase,
		}

		mgr.Responder.Respond(w, http.StatusCreated, inviteJSON)
	}
}

// GetInvite retrieves an invite based on the id given
func (mgr *InviteMgr) GetInvite() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		inviteId := mux.Vars(req)["inviteId"]

		invite, err := mgr.InviteStore.Get(inviteId)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		var qrcode invitesmgr.QRCode
		json.Unmarshal([]byte(invite.QRCode), &qrcode)

		inviteJSON := invitesmgr.InviteJSON{
			Id:         invite.Id,
			Organiser:  invite.Organiser,
			Location:   invite.Location,
			Date:       invite.Date,
			QRCode:     qrcode,
			Passphrase: invite.Passphrase,
		}

		mgr.Responder.Respond(w, http.StatusOK, inviteJSON)
	}
}

func (mgr *InviteMgr) Register() {
	mgr.Router.HandleFunc(ROOT, mgr.NewInvite()).Methods("POST")
	mgr.Router.HandleFunc(INVITE, mgr.GetInvite()).Methods("GET")
}