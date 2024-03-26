package invites

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitr.io.com/responder"
	invites_pkg "invitr.io.com/services/invites/endpoints/pkg"
	"invitr.io.com/services/invites/invitestore"
)

const (
	ROOT   = "/invites"
	BASE   = ROOT + "/invite"
	USER   = ROOT + "/user/{userId}"
	INVITE = BASE + "/{inviteId}"
)

type InviteMgr struct {
	Env         string
	Logger      *zap.Logger
	Router      *mux.Router
	Responder   responder.Responder
	InviteStore invitestore.InviteStorer
}

func NewInviteMgr(router *mux.Router, environment string, logger *zap.Logger, responder responder.Responder, invitestore invitestore.InviteStorer) *InviteMgr {
	e := &InviteMgr{
		Env:         environment,
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
		if mgr.Env != "dev" {
			_, err := invites_pkg.GetUser(getinvite.Organiser)
			if err != nil {
				mgr.Responder.Error(w, 404, err)
				return
			}
		}

		// check if the user has invited people
		if len(getinvite.Invitees) == 0 {
			mgr.Responder.Error(w, 401, errors.New("no-one was invited, you need to invite someone"))
			return
		}

		// we need to generate a QR code, for this we need to
		// call the QR code microservice
		qrcode, err := invites_pkg.GenerateQRCode()
		if err != nil {
			mgr.Responder.Error(w, 500, err)
			return
		}

		// we could possibly send it as a byte array
		// but we'll need to unmarshal it properly
		qrcodeBytes := invites_pkg.QrToBytes(*qrcode)
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
				Title:      invite.Title,
				Organiser:  invite.Organiser,
				Location:   invite.Location,
				Date:       invite.Date,
				Passphrase: invite.Passphrase,
				Invitees:   invite.Invitees,
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
		qrcode := invites_pkg.BytesToQR([]byte(invite.QRCode))
		inviteJSON := invitestore.InviteJSON{
			Invite: invitestore.Invite{
				Id:         invite.Id,
				Title:      invite.Title,
				Organiser:  invite.Organiser,
				Location:   invite.Location,
				Date:       invite.Date,
				Passphrase: invite.Passphrase,
				Invitees:   invite.Invitees,
			},
			QRCode: qrcode,
		}

		mgr.Responder.Respond(w, http.StatusOK, inviteJSON)
	}
}

// GetInvite retrieves an invite based on the id given
func (mgr *InviteMgr) GetInvitesByUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := mux.Vars(req)["userId"]

		invites, err := mgr.InviteStore.ListByUser(userId)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		invitesJSON := make([]invitestore.InviteJSON, 0)

		// need to change bytes to QR from db and then
		// return in the JSON format
		for _, invite := range invites {
			qrcode := invites_pkg.BytesToQR([]byte(invite.QRCode))
			invitesJSON = append(invitesJSON, invitestore.InviteJSON{
				Invite: invitestore.Invite{
					Id:         invite.Id,
					Title:      invite.Title,
					Organiser:  invite.Organiser,
					Location:   invite.Location,
					Date:       invite.Date,
					Passphrase: invite.Passphrase,
					Invitees:   invite.Invitees,
				},
				QRCode: qrcode,
			})

		}

		mgr.Responder.Respond(w, http.StatusOK, invitesJSON)
	}
}

func (mgr *InviteMgr) Register() {
	mgr.Router.HandleFunc(BASE, mgr.NewInvite()).Methods("POST")
	mgr.Router.HandleFunc(INVITE, mgr.GetInvite()).Methods("GET")
	mgr.Router.HandleFunc(USER, mgr.GetInvitesByUser()).Methods("GET")
}
