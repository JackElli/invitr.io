package endpoints

import (
	"database/sql"
	"users/endpoints/user"
	"users/responder"
	"users/usermgr"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

const database = "users"

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

	db, err := sql.Open("mysql", "todo:todosecret@tcp(db-users)/"+database)
	if err != nil {
		return err
	}

	// set up stores, this is where we interact with the
	// db
	userStore := usermgr.NewUserStore(e.Logger, db)

	public := r.PathPrefix("/").Subrouter()
	// add endpoints to the router
	_ = user.NewUserMgr(public, e.Logger, responder, userStore)

	return nil
}
