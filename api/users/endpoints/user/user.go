package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitio.com/users/responder"
	"invitio.com/users/userstore"
)

const (
	ROOT = "/user"
	USER = ROOT + "/{userId}"
)

type UserMgr struct {
	Logger    *zap.Logger
	Router    *mux.Router
	Responder responder.Responder
	UserStore userstore.UserStorer
}

func NewUserMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, userstore userstore.UserStorer) *UserMgr {
	e := &UserMgr{
		Logger:    logger,
		Router:    router,
		Responder: responder,
		UserStore: userstore,
	}
	e.Register()
	return e
}

// GetUser returns a user given a user id
func (mgr *UserMgr) GetUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := mux.Vars(req)["userId"]

		user, err := mgr.UserStore.Get(userId)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		if user == nil {
			mgr.Responder.Error(w, 404, errors.New("user not found"))
			return
		}

		mgr.Responder.Respond(w, http.StatusOK, user)
	}
}

// NewUser adds a new user to the db
func (mgr *UserMgr) NewUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var getuser userstore.User
		json.NewDecoder(req.Body).Decode(&getuser)

		user, err := mgr.UserStore.Insert(&getuser)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		mgr.Responder.Respond(w, http.StatusCreated, user)
	}
}

// RemoveUser removes a user from the db given a userId
func (mgr *UserMgr) RemoveUser() func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := mux.Vars(req)["userId"]

		user, err := mgr.UserStore.Get(userId)
		if err != nil {
			mgr.Responder.Error(w, 500, err)
			return
		}

		if user == nil {
			mgr.Responder.Error(w, 404, errors.New("no users exist with that id"))
			return
		}

		err = mgr.UserStore.Remove(userId)
		if err != nil {
			mgr.Responder.Error(w, 400, err)
			return
		}

		// maybe we could have a unified way of sending
		// responses back by using a struct
		// something like
		// {
		//   message: "hello"
		// }

		message := "Successfully removed user."
		mgr.Responder.Respond(w, http.StatusAccepted, message)
	}
}

func (mgr *UserMgr) Register() {
	mgr.Router.HandleFunc(USER, mgr.GetUser()).Methods("GET")
	mgr.Router.HandleFunc(USER, mgr.RemoveUser()).Methods("DELETE")
	mgr.Router.HandleFunc(ROOT, mgr.NewUser()).Methods("POST")
}
