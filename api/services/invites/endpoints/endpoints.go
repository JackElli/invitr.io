package endpoints

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"invitr.io.com/responder"
	"invitr.io.com/services/invites/endpoints/invites"
	"invitr.io.com/services/invites/invitestore"
)

const DB = "invites"

type Endpoints struct {
	Logger *zap.Logger
}

func NewEndpointsMgr(logger *zap.Logger) *Endpoints {
	return &Endpoints{
		Logger: logger,
	}
}

// getConnectionString creates the mariadb connection string based on
// the creds given
func getConnectionString(username string, password string, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/", username, password, dbname) + DB
}

// SetupEndpoints creates the store and routes for the api to run
func (e *Endpoints) SetupEndpoints(env string, r *mux.Router) error {
	responder := responder.NewResponder()

	connection := getConnectionString(
		os.Getenv("INVITES_DB_USERNAME"),
		os.Getenv("INVITES_DB_PASSWORD"),
		os.Getenv("INVITES_DB_NAME"),
	)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		return err
	}

	public := r.PathPrefix("/").Subrouter()

	// set up stores, this is where we interact with the db
	inviteStore := invitestore.NewInviteStore(e.Logger, db)

	// add endpoints to the router
	_ = invites.NewInviteMgr(public, env, e.Logger, responder, inviteStore)

	return nil
}
