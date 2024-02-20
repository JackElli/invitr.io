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
	Remove(userId string) error

	// Add these in later (not needed for now)
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
	row := us.db.QueryRow(query)

	var user User
	switch err := row.Scan(&user.UserId, &user.Username); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}

// Insert adds a user to the db
func (us *UserStore) Insert(user *User) (*User, error) {
	id, _ := uuid.NewV7()
	query := fmt.Sprintf("INSERT INTO users (id, username) VALUES ('%s','%s')", id, user.Username)

	_, err := us.db.Query(query)
	if err != nil {
		return nil, err
	}

	user.UserId = id.String()
	return user, err
}

// Remove removes a user from the db
func (us *UserStore) Remove(userId string) error {
	query := fmt.Sprintf("DELETE FROM users WHERE id='%s'", userId)

	_, err := us.db.Query(query)

	return err
}
