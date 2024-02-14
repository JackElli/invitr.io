package usermgr

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"
)

type UserStorer interface {
	Get(id string) (*User, error)
	// Query(querystr string, options *gocb.QueryOptions) ([]Grumble, error)
	// Insert(id string, grumble *Grumble) error
	// Update(id string, grumble *Grumble) error
}

type UserStore struct {
	Logger *zap.Logger
	db     *sql.DB
}

func NewUserStore(logger *zap.Logger, db *sql.DB) *UserStore {
	// need to create the table if it doesn't exist
	db.Query("CREATE TABLE IF NOT EXISTS users (id varchar(50) NOT NULL, username varchar(50) NOT NULL);")
	return &UserStore{
		Logger: logger,
		db:     db,
	}
}

func (us *UserStore) Get(id string) (*User, error) {
	query := fmt.Sprintf("SELECT id,username FROM users WHERE id='%s'", id)

	var userid string
	var username string

	// we could use a User struct here
	row := us.db.QueryRow(query)

	switch err := row.Scan(&userid, &username); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &User{
			UserId:   userid,
			Username: username,
		}, nil
	default:
		return nil, err
	}
}
