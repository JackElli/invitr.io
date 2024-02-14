package endpoints

import (
	"database/sql"
	"users/endpoints/user"
	"users/managers/usermgr"
	"users/responder"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
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

	db, err := sql.Open("mysql", "todo:todosecret@tcp(db-users)/users")
	if err != nil {
		return err
	}

	// set up managers, this is where we interact with the
	// db in the backend
	userStore := usermgr.NewUserStore(e.Logger, db)

	// add endpoints to the router
	public := r.PathPrefix("/").Subrouter()
	_ = user.NewUserMgr(public, e.Logger, responder, userStore)
	return nil
}
