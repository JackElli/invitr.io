package endpoints

import (
	"database/sql"
	"invites/endpoints/invite"
	"invites/invitestore"
	"invites/responder"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB   = "invites"
	CONN = "todo:todosecret@tcp(db-invites)/" + DB
)

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

	db, err := sql.Open("mysql", CONN)
	if err != nil {
		return err
	}

	// set up stores, this is where we interact with the
	// db
	inviteStore := invitestore.NewInviteStore(e.Logger, db)

	public := r.PathPrefix("/").Subrouter()
	// add endpoints to the router
	_ = invite.NewInviteMgr(public, e.Logger, responder, inviteStore)

	return nil
}
