package orgstore

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type OrgStorer interface {
	Get(id string) (*Organisation, error)
	Insert(invite *Organisation) (*Organisation, error)
	Update(table string, orgId string, fieldName string, fieldValue string) error
	Query(querystr string, args ...any) (*sql.Rows, error)
	// Remove(userId string) error
}

type OrgStore struct {
	Logger *zap.Logger
	db     *sql.DB
}

func NewOrgStore(logger *zap.Logger, db *sql.DB) *OrgStore {
	// need to create the table if it doesn't exist
	_, err := db.Query("CREATE TABLE IF NOT EXISTS organisations (id varchar(50) NOT NULL, name varchar(50) NOT NULL);")

	if err != nil {
		logger.Error(err.Error())
		// if we cannot connect to the db, we panic
		// and let docker restart it
		panic("Cannot create invites table.")
	}

	// will also need to create the pivot tables here too :)
	_, err = db.Query("CREATE TABLE IF NOT EXISTS organisation_people (id varchar(50) NOT NULL, organisation_id varchar(50) NOT NULL, user varchar(50) NOT NULL);")

	if err != nil {
		logger.Error(err.Error())
		// if we cannot connect to the db, we panic
		// and let docker restart it
		panic("Cannot create invites_invitee table.")
	}

	logger.Info("Created all tables.")
	return &OrgStore{
		Logger: logger,
		db:     db,
	}
}

func (store *OrgStore) InitDemoOrg() {

	demoOrg := Organisation{
		Name: "Jacks Test Org",
		People: []string{
			"123",
			"234",
		},
	}

	store.Insert(&demoOrg)
}

// Get retrieves in this case a user from the db
func (store *OrgStore) Get(id string) (*Organisation, error) {
	// join the invite and invitees table because 1 invite can have many invitees
	// we then need to aggregate the invitees into a json list. Once we have the
	// query response, the invitees will be in bytes so we need to json unmarshal
	// to turn it into a string list.
	row := store.db.QueryRow("SELECT o.id, o.name, IF(op.user IS NOT NULL, JSON_ARRAYAGG(op.user), NULL) as people FROM organisations o LEFT JOIN organisation_people op ON op.organisation_id=o.id WHERE o.id = ?", id)

	var org Organisation
	var people []byte
	switch err := row.Scan(&org.Id, &org.Name, &people); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		json.Unmarshal(people, &org.People)
		return &org, nil
	default:
		return nil, err
	}
}

// Insert adds an invite to the db
func (store *OrgStore) Insert(org *Organisation) (*Organisation, error) {
	id, _ := uuid.NewV7()

	// TODO for the QR code, can we use a byte array instead of JSON string
	// we will probably need to use %v for the format.
	_, err := store.db.Query("INSERT INTO organisations (id, name) VALUES (?, ?)", id, org.Name)
	if err != nil {
		return nil, err
	}

	// need to add the people to the pivot table
	for _, userId := range org.People {
		orgPeopleId, _ := uuid.NewV7()
		_, err := store.db.Query("INSERT INTO organisation_people (id, organisation_id, user) VALUES (?, ?, ?)", &orgPeopleId, id, userId)
		if err != nil {
			return nil, err
		}
	}

	org.Id = id.String()
	return org, err
}

// Update updates a field within the invites table
// TODO make this work with other data types
func (store *OrgStore) Update(table string, inviteId string, fieldName string, fieldValue string) error {
	_, err := store.db.Query("UPDATE ? SET ? = ? WHERE id = ?", table, fieldName, fieldValue, inviteId)

	return err
}

// Query runs a custom query on the DB (need to be careful here :) )
func (store *OrgStore) Query(querystr string, args ...any) (*sql.Rows, error) {
	rows, err := store.db.Query(querystr, args)
	return rows, err
}
