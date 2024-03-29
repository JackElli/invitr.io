package invites

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitr.io.com/responder"
	invites_pkg "invitr.io.com/services/invites/endpoints/pkg"
	"invitr.io.com/services/invites/invitestore"
)

const (
	ROOT    = "/invites"
	USER    = ROOT + "/user/{userId}"
	BASE    = ROOT + "/invite"
	INVITE  = BASE + "/{inviteId}"
	NOTE    = INVITE + "/note"
	EVENT   = INVITE + "/user/{user}"
	ORG_KEY = EVENT + "/key"
	KEY     = INVITE + "/key"
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

type InviteNote struct {
	Notes string `json:"notes"`
}

// AddNotes adds notes to a given invite based on its ID
func (mgr *InviteMgr) AddNotes() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		inviteId := mux.Vars(req)["inviteId"]

		var getnote InviteNote
		json.NewDecoder(req.Body).Decode(&getnote)

		err := mgr.InviteStore.Update("invites", inviteId, "notes", getnote.Notes)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		mgr.Responder.Respond(w, http.StatusCreated, "Successfully updated notes")
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
				Notes:      invite.Notes,
				Date:       invite.Date,
				Passphrase: invite.Passphrase,
				Invitees:   invite.Invitees,
			},
			QRCode: qrcode,
		}

		mgr.Responder.Respond(w, http.StatusOK, inviteJSON)
	}
}

// ListInvitesByUser retrieves an invite based on the user given
func (mgr *InviteMgr) ListInvitesByUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := mux.Vars(req)["userId"]

		invites, err := mgr.InviteStore.ListByUser(userId)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		invitesJSON := make(map[string][]invitestore.InviteJSON, 0)

		for _, invite := range invites {
			// need to change bytes to QR from db and then
			qrcode := invites_pkg.BytesToQR([]byte(invite.QRCode))
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
				QRCode: qrcode,
			}

			// check if date of event has passed
			date, _ := time.Parse("2006-01-02 15:04:05", invite.Date)
			if date.After(time.Now()) {
				invitesJSON["ongoing"] = append(invitesJSON["ongoing"], inviteJSON)
			} else {
				invitesJSON["finished"] = append(invitesJSON["finished"], inviteJSON)
			}
		}

		mgr.Responder.Respond(w, http.StatusOK, invitesJSON)
	}
}

// GetInvite retrieves an invite based on the id given
func (mgr *InviteMgr) IsUserGoingToEvent() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		inviteId := mux.Vars(req)["inviteId"]
		user := mux.Vars(req)["user"]

		// TODO could we make the table names const?
		rows, err := mgr.InviteStore.Query(
			fmt.Sprintf("SELECT is_going FROM invites_invitees WHERE invite_id='%s' AND invitee='%s'",
				inviteId, user),
		)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		var going *bool
		for rows.Next() {
			rows.Scan(&going)
		}

		mgr.Responder.Respond(w, http.StatusOK, going)
	}
}

type EventResponse struct {
	Going bool `json:"going"`
}

// GetInvite retrieves an invite based on the id given
func (mgr *InviteMgr) RespondToEvent() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		inviteId := mux.Vars(req)["inviteId"]
		user := mux.Vars(req)["user"]

		var eventresp EventResponse
		json.NewDecoder(req.Body).Decode(&eventresp)

		// TODO can we find a more elegant way of doing this?
		going := 0
		if eventresp.Going {
			going = 1
		}

		// TODO could we make the table names const?
		_, err := mgr.InviteStore.Query(
			fmt.Sprintf("UPDATE invites_invitees SET is_going='%d' WHERE invite_id='%s' AND invitee='%s'",
				going, inviteId, user),
		)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		mgr.Responder.Respond(w, http.StatusOK, "Successfully updated going field")
	}
}

type UserKey struct {
	Key string `json:"key"`
}

// GetUserFromKey retrieves an invite based on the id given
func (mgr *InviteMgr) GetUserFromKey() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		inviteId := mux.Vars(req)["inviteId"]

		var userkey UserKey
		json.NewDecoder(req.Body).Decode(&userkey)

		rows, err := mgr.InviteStore.Query(
			fmt.Sprintf("SELECT invitee FROM invites_invitees WHERE invite_id='%s' AND invite_key='%s'",
				inviteId, userkey.Key),
		)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		var user *string
		for rows.Next() {
			rows.Scan(&user)
		}

		mgr.Responder.Respond(w, http.StatusOK, user)
	}
}

// WARNING DANGER!!! This needs to be locked down!!! for now it's exposed
// BUT WE NEED TO LOCK IT DOWN EVENTUALLY
// MAYBE MAKE IT SO YOU NEED ANOTHER KEY TO GET THE KEY
func (mgr *InviteMgr) GetOrganiserKey() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		inviteId := mux.Vars(req)["inviteId"]
		user := mux.Vars(req)["user"]

		rows, err := mgr.InviteStore.Query(
			fmt.Sprintf("SELECT invite_key FROM invites_invitees WHERE invite_id='%s' AND invitee='%s'",
				inviteId, user),
		)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		var key *string
		for rows.Next() {
			rows.Scan(&key)
		}

		mgr.Responder.Respond(w, http.StatusOK, key)
	}
}

func (mgr *InviteMgr) Register() {
	mgr.Router.HandleFunc(BASE, mgr.NewInvite()).Methods("POST")
	mgr.Router.HandleFunc(NOTE, mgr.AddNotes()).Methods("POST")
	mgr.Router.HandleFunc(EVENT, mgr.RespondToEvent()).Methods("POST")
	mgr.Router.HandleFunc(KEY, mgr.GetUserFromKey()).Methods("POST")
	// DANGER PLEASE LOCK THIS DOWN!!!
	mgr.Router.HandleFunc(ORG_KEY, mgr.GetOrganiserKey()).Methods("GET")

	mgr.Router.HandleFunc(INVITE, mgr.GetInvite()).Methods("GET")
	mgr.Router.HandleFunc(USER, mgr.ListInvitesByUser()).Methods("GET")
	mgr.Router.HandleFunc(EVENT, mgr.IsUserGoingToEvent()).Methods("GET")
}
