package endpoints

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitr.io.com/invites/endpoints/invite"
	"invitr.io.com/invites/invitestore"
	"invitr.io.com/responder"
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
	_ = invite.NewInviteMgr(public, env, e.Logger, responder, inviteStore)

	return nil
}
