package usermgr

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserStorer interface {
	Get(id string) (*User, error)
	Insert(user *User) (*User, error)
	// Query(querystr string, options *gocb.QueryOptions) ([]Grumble, error)
	// Update(id string, grumble *Grumble) error
}

type UserStore struct {
	Logger *zap.Logger
	db     *sql.DB
}

func NewUserStore(logger *zap.Logger, db *sql.DB) *UserStore {
	// need to create the table if it doesn't exist
	_, err := db.Query("CREATE TABLE IF NOT EXISTS users (id varchar(50) NOT NULL, username varchar(50) NOT NULL);")
	if err != nil {
		logger.Error(err.Error())
		// if we cannot connect to the db, we panic
		// and let docker restart it
		panic("Cannot create table.")
	}

	logger.Info("Created users table.")
	return &UserStore{
		Logger: logger,
		db:     db,
	}
}

// Get retrieves in this case a user from the db
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

// Insert adds a user to the db
func (us *UserStore) Insert(user *User) (*User, error) {
	id, _ := uuid.NewV7()
	query := fmt.Sprintf("INSERT INTO users (id, username) VALUES ('%s','%s')", id, user.Username)

	_, err := us.db.Query(query)

	var rUser User
	rUser.UserId = id.String()
	rUser.Username = user.Username

	return &rUser, err
}
