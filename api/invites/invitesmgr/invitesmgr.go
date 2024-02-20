package invitesmgr

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type InviteStorer interface {
	Get(id string) (*Invite, error)
	Insert(user *Invite) (*Invite, error)
	Remove(userId string) error

	// Add these in later (not needed for now)
	// Query(querystr string, options *gocb.QueryOptions) ([]Grumble, error)
	// Update(id string, grumble *Grumble) error
}

type InviteStore struct {
	Logger *zap.Logger
	db     *sql.DB
}

func NewInviteStore(logger *zap.Logger, db *sql.DB) *InviteStore {
	// need to create the table if it doesn't exist
	_, err := db.Query("CREATE TABLE IF NOT EXISTS invites (id varchar(50) NOT NULL, username varchar(50) NOT NULL);")
	if err != nil {
		logger.Error(err.Error())
		// if we cannot connect to the db, we panic
		// and let docker restart it
		panic("Cannot create table.")
	}

	logger.Info("Created invites table.")
	return &InviteStore{
		Logger: logger,
		db:     db,
	}
}

// Get retrieves in this case a user from the db
func (us *InviteStore) Get(id string) (*Invite, error) {
	query := fmt.Sprintf("SELECT id,username FROM users WHERE id='%s'", id)

	var userid string
	var username string

	// we could use a User struct here
	row := us.db.QueryRow(query)

	switch err := row.Scan(&userid, &username); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &Invite{
			UserId:   userid,
			Username: username,
		}, nil
	default:
		return nil, err
	}
}
