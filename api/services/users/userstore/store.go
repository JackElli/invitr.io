package userstore

import (
	"database/sql"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserStorer interface {
	Get(id string) (*User, error)
	Insert(user *User) (*User, error)
	Remove(userId string) error

	// Add these in later (not needed for now)
	// Query(querystr string, args ...any) (*sql.Rows, error)
}

type UserStore struct {
	Logger *zap.Logger
	db     *sql.DB
}

func NewUserStore(logger *zap.Logger, db *sql.DB) *UserStore {
	// need to create the table if it doesn't exist
	_, err := db.Query("CREATE TABLE IF NOT EXISTS users (id varchar(50) NOT NULL, first_name varchar(50) NOT NULL, last_name varchar(50) NOT NULL);")
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

func (store *UserStore) InitDemoUser() {
	demoUser := User{
		Id:        "123",
		FirstName: "Demo",
		LastName:  "User",
	}

	store.Insert(&demoUser)
}

// Get retrieves in this case a user from the db
func (us *UserStore) Get(id string) (*User, error) {
	row := us.db.QueryRow("SELECT id, first_name, last_name FROM users WHERE id = ?", id)

	var user User
	switch err := row.Scan(&user.Id, &user.FirstName, &user.LastName); err {
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
	// if it's not a demo user, we need to generate
	// and ID
	if user.Id == "" {
		id, _ := uuid.NewV7()
		user.Id = id.String()
	}

	_, err := us.db.Query("INSERT INTO users (id, first_name, last_name) VALUES (?, ?, ?)", user.Id, user.FirstName, user.LastName)
	if err != nil {
		return nil, err
	}

	return user, err
}

// Remove removes a user from the db
func (us *UserStore) Remove(userId string) error {
	_, err := us.db.Query("DELETE FROM users WHERE id = ?", userId)

	return err
}
