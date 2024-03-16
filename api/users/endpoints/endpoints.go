package endpoints

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitr.io.com/responder"
	"invitr.io.com/users/endpoints/user"
	"invitr.io.com/users/userstore"
)

const (
	DB   = "users"
	CONN = "todo:todosecret@tcp(db-users)/" + DB
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
	userStore := userstore.NewUserStore(e.Logger, db)

	public := r.PathPrefix("/").Subrouter()
	// add endpoints to the router
	_ = user.NewUserMgr(public, e.Logger, responder, userStore)

	return nil
}
