package invites

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitr.io.com/responder"
	invites_fetch "invitr.io.com/services/invites/fetch"
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
	QRMgr       invites_fetch.IQRMgr
}

func NewInviteMgr(router *mux.Router, environment string, logger *zap.Logger, responder responder.Responder, invitestore invitestore.InviteStorer, qrmgr invites_fetch.IQRMgr) *InviteMgr {
	e := &InviteMgr{
		Env:         environment,
		Logger:      logger,
		Router:      router,
		Responder:   responder,
		InviteStore: invitestore,
		QRMgr:       qrmgr,
	}
	e.Register()
	return e
}

// NewInvite creates a new invite based on some user input
func (mgr *InviteMgr) New() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var getinvite invitestore.Invite
		json.NewDecoder(req.Body).Decode(&getinvite)

		// we need to check if the organiser is actually
		// a registered user
		if mgr.Env != "dev" {
			_, err := invites_fetch.GetUser(getinvite.Organiser)
			if err != nil {
				mgr.Responder.Error(w, 404, err)
				return
			}
		}

		// check if the user has invited people
		if len(getinvite.Invitees) == 0 {
			err := errors.New("no-one was invited, you need to invite someone")
			mgr.Logger.Error(err.Error())
			mgr.Responder.Error(w, 400, err)
			return
		}

		// we need to generate a QR code, for this we need to
		// call the QR code microservice
		qrcode, err := mgr.QRMgr.GenerateQRCode()
		if err != nil {
			mgr.Logger.Error(err.Error())
			mgr.Responder.Error(w, 500, err)
			return
		}

		// we could possibly send it as a byte array
		// but we'll need to unmarshal it properly
		qrcodeBytes := mgr.QRMgr.QrToBytes(*qrcode)
		inviteDB := invitestore.InviteDB{
			Invite: getinvite,
			QRCode: string(qrcodeBytes),
		}

		invite, err := mgr.InviteStore.Insert(&inviteDB)
		if err != nil {
			mgr.Logger.Error(err.Error())
			mgr.Responder.Error(w, 400, err)
			return
		}

		// return in the JSON format
		inviteJSON := invitestore.InviteJSON{
			Invite: invite.Invite,
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

		// TODO add check to see if note is valid
		// TODO have a think if this needs to be this generic?
		err := mgr.InviteStore.Update("invites", inviteId, "notes", getnote.Notes)
		if err != nil {
			mgr.Logger.Error(err.Error())
			mgr.Responder.Error(w, 400, err)
			return
		}

		mgr.Responder.Respond(w, http.StatusCreated, "Successfully updated notes")
	}
}

// GetInvite retrieves an invite based on the id given
func (mgr *InviteMgr) Get() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		inviteId := mux.Vars(req)["inviteId"]

		invite, err := mgr.InviteStore.Get(inviteId)
		if err != nil {
			mgr.Logger.Error(err.Error())
			mgr.Responder.Error(w, 400, err)
			return
		}

		if invite == nil {
			mgr.Responder.Error(w, 404, errors.New("invite not found"))
			return
		}

		// need to change bytes to QR from db and then
		// return in the JSON format
		qrcode := mgr.QRMgr.BytesToQR([]byte(invite.QRCode))
		inviteJSON := invitestore.InviteJSON{
			// can be dereferenced because we already checked to
			// see if it's nil
			Invite: invitestore.InviteDBtoInvite(*invite),
			QRCode: qrcode,
		}

		mgr.Responder.Respond(w, http.StatusOK, inviteJSON)
	}
}

// ListInvitesByUser retrieves an invite based on the user given
func (mgr *InviteMgr) ListByUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := mux.Vars(req)["userId"]

		invites, err := mgr.InviteStore.ListByUser(userId)
		if err != nil {
			mgr.Logger.Error(err.Error())
			mgr.Responder.Error(w, 400, err)
			return
		}

		invitesJSON := make(map[string][]invitestore.InviteJSON, 0)

		for _, invite := range invites {
			// need to change bytes to QR from db and then
			qrcode := mgr.QRMgr.BytesToQR([]byte(invite.QRCode))
			// return in the JSON format
			inviteJSON := invitestore.InviteJSON{
				Invite: invitestore.InviteDBtoInvite(invite),
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
			"SELECT is_going FROM invites_invitees WHERE invite_id = ? AND invitee = ?",
			inviteId, user)
		if err != nil {
			mgr.Logger.Error(err.Error())
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
			"UPDATE invites_invitees SET is_going = ?  WHERE invite_id = ? AND invitee = ?",
			going, inviteId, user,
		)
		if err != nil {
			mgr.Logger.Error(err.Error())
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
			"SELECT invitee FROM invites_invitees WHERE invite_id = ? AND invite_key = ?",
			inviteId, userkey.Key,
		)
		if err != nil {
			mgr.Logger.Error(err.Error())
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
			"SELECT invite_key FROM invites_invitees WHERE invite_id = ? AND invitee = ?",
			inviteId, user,
		)
		if err != nil {
			mgr.Logger.Error(err.Error())
			mgr.Responder.Error(w, 400, err)
			return
		}

		var key string
		for rows.Next() {
			rows.Scan(&key)
		}

		mgr.Responder.Respond(w, http.StatusOK, key)
	}
}

func (mgr *InviteMgr) Register() {
	mgr.Router.HandleFunc(BASE, mgr.New()).Methods("POST")
	mgr.Router.HandleFunc(NOTE, mgr.AddNotes()).Methods("POST")
	mgr.Router.HandleFunc(EVENT, mgr.RespondToEvent()).Methods("POST")
	mgr.Router.HandleFunc(KEY, mgr.GetUserFromKey()).Methods("POST")

	/////DANGER PLEASE LOCK THIS DOWN!!!////
	mgr.Router.HandleFunc(ORG_KEY, mgr.GetOrganiserKey()).Methods("GET")
	////////////////////////////////////////

	mgr.Router.HandleFunc(INVITE, mgr.Get()).Methods("GET")
	mgr.Router.HandleFunc(USER, mgr.ListByUser()).Methods("GET")
	mgr.Router.HandleFunc(EVENT, mgr.IsUserGoingToEvent()).Methods("GET")
}
