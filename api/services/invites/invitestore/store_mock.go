package invitestore

import (
	"database/sql"
)

type InviteStoreMock struct{}

func NewInviteStoreMock() *InviteStoreMock {
	return &InviteStoreMock{}
}

func (store *InviteStoreMock) Get(id string) (*InviteDB, error) {
	if id == "1234" {
		return &InviteDB{
			Invite: Invite{
				Id:    "1234",
				Title: "Hello",
			},
		}, nil
	}

	return nil, nil
}

func (store *InviteStoreMock) ListByUser(userId string) ([]InviteDB, error) {
	return nil, nil
}

func (store *InviteStoreMock) Insert(invite *InviteDB) (*InviteDB, error) {
	var _notGoing int8 = 0
	notGoing := &_notGoing

	return &InviteDB{
		Invite: Invite{
			Id:    "1234",
			Title: "Hello",
			Invitees: []Person{{
				Id:      "12345",
				IsGoing: notGoing,
			}},
		},
	}, nil
}

func (store *InviteStoreMock) Update(table string, inviteId string, fieldName string, fieldValue string) error {
	return nil
}

func (store *InviteStoreMock) Query(querystr string) (*sql.Rows, error) {
	return nil, nil
}
