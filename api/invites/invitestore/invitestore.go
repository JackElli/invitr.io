package invitestore

import (
	"database/sql"
	"encoding/json"
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

	if err != nil {
		logger.Error(err.Error())
		// if we cannot connect to the db, we panic
		// and let docker restart it
		panic("Cannot create invites table.")
	}

	// will also need to create the pivot tables here too :)
	_, err = db.Query("CREATE TABLE IF NOT EXISTS invites_invitees (id varchar(50) NOT NULL, invite_id varchar(50) NOT NULL, invitee varchar(50) NOT NULL);")

	if err != nil {
		logger.Error(err.Error())
		// if we cannot connect to the db, we panic
		// and let docker restart it
		panic("Cannot create invites_invitee table.")
	}

	logger.Info("Created all tables.")
	return &InviteStore{
		Logger: logger,
		db:     db,
	}
}

// Get retrieves in this case a user from the db
func (is *InviteStore) Get(id string) (*InviteDB, error) {
	// join the invite and invitees table because 1 invite can have many invitees
	// we then need to aggregate the invitees into a json list. Once we have the
	// query response, the invitees will be in bytes so we need to json unmarshal
	// to turn it into a string list.
	query := fmt.Sprintf("SELECT i.id, i.organiser, i.location, i.date, i.qr_code, i.passphrase, IF(ii.invitee IS NOT NULL, JSON_ARRAYAGG(ii.invitee), NULL) as invitees FROM invites i LEFT JOIN invites_invitees ii ON ii.invite_id=i.id WHERE i.id='%s'", id)
	row := is.db.QueryRow(query)

	var invite InviteDB
	var invitees []byte
	switch err := row.Scan(&invite.Id,
		&invite.Organiser, &invite.Location,
		&invite.Date, &invite.QRCode, &invite.Passphrase, &invitees); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		// need to JSONify the invitees
		json.Unmarshal(invitees, &invite.Invitees)
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

	// need to add the invitees to the pivot table
	inviteInviteeId, _ := uuid.NewV7()
	for _, invitee := range invite.Invitees {
		query := fmt.Sprintf("INSERT INTO invites_invitees (id, invite_id, invitee) VALUES ('%s','%s', '%s')", inviteInviteeId, id, string(invitee))

		_, err := is.db.Query(query)
		if err != nil {
			return nil, err
		}
	}

	invite.Id = id.String()
	return invite, err
}
