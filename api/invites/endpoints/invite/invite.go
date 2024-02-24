package invite

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitio.com/invites/invitestore"
	"invitio.com/invites/responder"
)

const (
	ROOT   = "/invite"
	INVITE = ROOT + "/{inviteId}"
)

type InviteMgr struct {
	Logger      *zap.Logger
	Router      *mux.Router
	Responder   responder.Responder
	InviteStore invitestore.InviteStorer
}

func NewInviteMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, invitestore invitestore.InviteStorer) *InviteMgr {
	e := &InviteMgr{
		Logger:      logger,
		Router:      router,
		Responder:   responder,
		InviteStore: invitestore,
	}
	e.Register()
	return e
}

// NewInvite creates a new invite based on some user input
func (mgr *InviteMgr) NewInvite() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var getinvite invitestore.InviteDB
		json.NewDecoder(req.Body).Decode(&getinvite)

		// we need to check if the organiser is actually
		// a registered user
		_, err := mgr.GetUser(getinvite.Organiser)
		if err != nil {
			mgr.Responder.Error(w, 404, err)
			return
		}

		// we need to generate a QR code, for this we need to
		// call the QR code microservice
		qrcode, err := GenerateQRCode()
		if err != nil {
			mgr.Responder.Error(w, 500, err)
			return
		}

		// We could possibly send it as a byte array
		// but we'll need to unmarshal it properly
		qrcodeBytes := qrToBytes(*qrcode)
		getinvite.QRCode = string(qrcodeBytes)
		invite, err := mgr.InviteStore.Insert(&getinvite)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		// return in the JSON format
		inviteJSON := invitestore.InviteJSON{
			Invite: invitestore.Invite{
				Id:         invite.Id,
				Organiser:  invite.Organiser,
				Location:   invite.Location,
				Date:       invite.Date,
				Passphrase: invite.Passphrase,
			},
			QRCode: *qrcode,
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

		if invite == nil {
			mgr.Responder.Respond(w, http.StatusOK, nil)
			return
		}

		// need to change bytes to QR from db and then
		// return in the JSON format
		qrcode := bytesToQR([]byte(invite.QRCode))
		inviteJSON := invitestore.InviteJSON{
			Invite: invitestore.Invite{
				Id:         invite.Id,
				Organiser:  invite.Organiser,
				Location:   invite.Location,
				Date:       invite.Date,
				Passphrase: invite.Passphrase,
			},
			QRCode: qrcode,
		}

		mgr.Responder.Respond(w, http.StatusOK, inviteJSON)
	}
}

func (mgr *InviteMgr) Register() {
	mgr.Router.HandleFunc(ROOT, mgr.NewInvite()).Methods("POST")
	mgr.Router.HandleFunc(INVITE, mgr.GetInvite()).Methods("GET")
}
