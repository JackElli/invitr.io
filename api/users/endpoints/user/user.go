package user

import (
	"net/http"
	"users/managers/usermgr"
	"users/responder"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

const (
	ROOT     = "/user"
	GET_USER = ROOT + "/{userId}"
)

type UserMgr struct {
	Logger    *zap.Logger
	Router    *mux.Router
	Responder responder.Responder
	UserStore usermgr.UserStorer
}

func NewUserMgr(router *mux.Router, logger *zap.Logger, responder responder.Responder, userstore usermgr.UserStorer) *UserMgr {
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

		// could we return a 404 if user not found?
		mgr.Responder.Respond(w, http.StatusOK, user)
	}
}

func (mgr *UserMgr) Register() {
	mgr.Router.HandleFunc(GET_USER, mgr.GetUser()).Methods("GET")
}
