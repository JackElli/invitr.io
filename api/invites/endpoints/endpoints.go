package endpoints

import (
	"database/sql"
	"invites/endpoints/invite"
	"invites/invitesmgr"
	"invites/responder"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

const database = "invites"

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

	db, err := sql.Open("mysql", "todo:todosecret@tcp(db-invites)/"+database)
	if err != nil {
		return err
	}

	// set up stores, this is where we interact with the
	// db
	inviteStore := invitesmgr.NewInviteStore(e.Logger, db)

	public := r.PathPrefix("/").Subrouter()
	// add endpoints to the router
	_ = invite.NewInviteMgr(public, e.Logger, responder, inviteStore)

	return nil
}
