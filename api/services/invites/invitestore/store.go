package invitestore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type InviteStorer interface {
	Get(id string) (*InviteDB, error)
	ListByUser(userId string) ([]InviteDB, error)
	Insert(invite *InviteDB) (*InviteDB, error)
	Update(table string, inviteId string, fieldName string, fieldValue string) error
	Query(querystr string, args ...any) (*sql.Rows, error)
	// Remove(userId string) error
}

type InviteStore struct {
	Logger *zap.Logger
	db     *sql.DB
}

func NewInviteStore(logger *zap.Logger, db *sql.DB) *InviteStore {
	// need to create the table if it doesn't exist
	_, err := db.Query("CREATE TABLE IF NOT EXISTS invites (id varchar(50) NOT NULL, title varchar(50) NOT NULL, organiser varchar(50) NOT NULL, location varchar(50) NOT NULL, notes varchar(255), date DATETIME NOT NULL, qr_code varchar(1024) NOT NULL, passphrase varchar(50) NOT NULL);")

	if err != nil {
		logger.Error(err.Error())
		// if we cannot connect to the db, we panic
		// and let docker restart it
		panic("Cannot create invites table.")
	}

	// will also need to create the pivot tables here too :)
	_, err = db.Query("CREATE TABLE IF NOT EXISTS invites_invitees (id varchar(50) NOT NULL, invite_id varchar(50) NOT NULL, invitee varchar(50) NOT NULL, is_going boolean, message varchar(100), invite_key varchar(50) NOT NULL);")

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
func (store *InviteStore) Get(id string) (*InviteDB, error) {
	// join the invite and invitees table because 1 invite can have many invitees
	// we then need to aggregate the invitees into a json list. Once we have the
	// query response, the invitees will be in bytes so we need to json unmarshal
	// to turn it into a string list.
	row := store.db.QueryRow("SELECT i.id, i.title, i.organiser, i.location, i.notes, i.date, i.qr_code, i.passphrase, IF(ii.invitee IS NOT NULL, JSON_ARRAYAGG(JSON_OBJECT('id', ii.invitee, 'is_going', ii.is_going)), NULL) as invitees FROM invites i LEFT JOIN invites_invitees ii ON ii.invite_id=i.id WHERE i.id = ?", id)

	var invite InviteDB
	var invitees []byte
	switch err := row.Scan(&invite.Id, &invite.Title, &invite.Organiser,
		&invite.Location, &invite.Notes, &invite.Date, &invite.QRCode, &invite.Passphrase,
		&invitees); err {
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

// GetByUser retrieves all invites created by a user
func (store *InviteStore) ListByUser(userId string) ([]InviteDB, error) {
	rows, err := store.db.Query("SELECT i.id, i.title, i.organiser, i.location, i.date, i.qr_code, i.passphrase, IF(ii.invitee IS NOT NULL, JSON_ARRAYAGG(ii.invitee), NULL) as invitees FROM invites i LEFT JOIN invites_invitees ii ON ii.invite_id=i.id WHERE ? IN (SELECT invitee FROM invites_invitees WHERE invite_id=i.id) GROUP BY i.id ORDER BY i.date ASC", userId)
	if err != nil {
		return nil, err
	}

	invites := make([]InviteDB, 0)

	var invite InviteDB
	var invitees []byte
	for rows.Next() {
		err := rows.Scan(&invite.Id, &invite.Title, &invite.Organiser,
			&invite.Location, &invite.Date, &invite.QRCode, &invite.Passphrase,
			&invitees)

		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(invitees, &invite.Invitees)

		invites = append(invites, invite)
	}

	return invites, nil
}

// Insert adds an invite to the db
func (store *InviteStore) Insert(invite *InviteDB) (*InviteDB, error) {
	id, _ := uuid.NewV7()

	// TODO for the QR code, can we use a byte array instead of JSON string
	// we will probably need to use %v for the format.
	_, err := store.db.Query("INSERT INTO invites (id, title, organiser, location, notes, date, qr_code, passphrase) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", id, invite.Title, invite.Organiser, invite.Location, invite.Notes, invite.Date, invite.QRCode, invite.Passphrase)
	if err != nil {
		return nil, err
	}

	// this is janky as anything
	var going_ int8 = 1
	going := &going_

	// prepend the organiser to the invitees
	invite.Invitees = append([]Person{
		{
			Id:      invite.Organiser,
			IsGoing: going,
		},
	}, invite.Invitees...)

	// need to add the invitees to the pivot table
	for _, invitee := range invite.Invitees {
		inviteInviteeId, _ := uuid.NewV7()
		inviteKey, _ := uuid.NewV7()

		_, err := store.db.Query("INSERT INTO invites_invitees (id, invite_id, invitee, invite_key) VALUES (?, ?, ?, ?)", inviteInviteeId, id, string(invitee.Id), inviteKey)
		if err != nil {
			return nil, err
		}
	}

	invite.Id = id.String()
	return invite, err
}

// Update updates a field within the invites table
// TODO make this work with other data types
func (store *InviteStore) Update(table string, inviteId string, fieldName string, fieldValue string) error {
	// can be sprintf here because table and fieldName are not user inputs
	_, err := store.db.Query(
		fmt.Sprintf("UPDATE %s SET %s = ? WHERE id = ?", table, fieldName),
		fieldValue, inviteId)

	return err
}

// Query runs a custom query on the DB (need to be careful here :) )
func (store *InviteStore) Query(querystr string, args ...any) (*sql.Rows, error) {

	// prepare the statement first, removing SQL injection
	// potential
	stmt, _ := store.db.Prepare(querystr)

	// then we query
	rows, err := stmt.Query(args...)
	return rows, err
}
