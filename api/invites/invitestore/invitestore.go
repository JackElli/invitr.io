package invitestore

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type InviteStorer interface {
	Get(id string) (*InviteDB, error)
	Insert(invite *InviteDB) (*InviteDB, error)
	// Remove(userId string) error

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
	_, err := db.Query("CREATE TABLE IF NOT EXISTS invites (id varchar(50) NOT NULL, organiser varchar(50) NOT NULL, location varchar(50) NOT NULL, date varchar(50) NOT NULL, qr_code varchar(1024) NOT NULL, passphrase varchar(50) NOT NULL);")

	// will also need to create the pivot tables here too :)
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
func (is *InviteStore) Get(id string) (*InviteDB, error) {
	query := fmt.Sprintf("SELECT * FROM invites WHERE id='%s'", id)
	row := is.db.QueryRow(query)

	var invite InviteDB
	switch err := row.Scan(&invite.Id,
		&invite.Organiser, &invite.Location,
		&invite.Date, &invite.QRCode, &invite.Passphrase); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &invite, nil
	default:
		return nil, err
	}
}

// Insert adds an invite to the db
func (is *InviteStore) Insert(invite *InviteDB) (*InviteDB, error) {
	id, _ := uuid.NewV7()

	// TODO for the QR code, can we use a byte array instead of JSON string
	// we will probably need to use %v for the format.
	query := fmt.Sprintf("INSERT INTO invites (id, organiser, location, date, qr_code, passphrase) VALUES ('%s','%s', '%s', '%s', '%s', '%s')", id, invite.Organiser, invite.Location, invite.Date, invite.QRCode, invite.Passphrase)

	_, err := is.db.Query(query)
	if err != nil {
		return nil, err
	}

	invite.Id = id.String()
	return invite, err
}
