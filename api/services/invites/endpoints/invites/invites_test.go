package invites

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gotest.tools/assert"
	"invitr.io.com/responder"

	invites_fetch "invitr.io.com/services/invites/fetch"
	"invitr.io.com/services/invites/invitestore"
	qrcodes "invitr.io.com/services/qr-codes/endpoints/qr-codes"
)

func TestNewInvite(t *testing.T) {
	type testcase struct {
		desc             string
		expectedResponse invitestore.InviteDB
		expectedStatus   int
	}

	responderMock := responder.NewResponder()
	invitestoreMock := invitestore.NewInviteStoreMock()
	loggerMock, _ := zap.NewProduction()
	qrMgrMock := invites_fetch.NewQRMgrMock()
	defer loggerMock.Sync()

	var _notGoing int8 = 0
	notGoing := &_notGoing

	newInvite := invitestore.InviteDB{
		Invite: invitestore.Invite{
			Id:    "1234",
			Title: "Hello",
			Invitees: []invitestore.Person{{
				Id:      "12345",
				IsGoing: notGoing,
			}},
		},
	}

	testcases := []testcase{
		{
			desc:             "HAPPY added invite correctly",
			expectedResponse: newInvite,
			expectedStatus:   201,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			newUserMgrMock := NewInviteMgr(
				rMock,
				"dev",
				loggerMock,
				responderMock,
				invitestoreMock,
				qrMgrMock,
			)
			newUserMgrMock.Register()

			w := httptest.NewRecorder()

			newInviteData, _ := json.Marshal(newInvite)

			r, _ := http.NewRequest("POST", BASE, bytes.NewBuffer(newInviteData))
			newUserMgrMock.Router.ServeHTTP(w, r)

			var response invitestore.InviteDB
			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
			assert.DeepEqual(t, response, testCase.expectedResponse)
		})
	}
}

func TestGetInvite(t *testing.T) {

	type testcase struct {
		desc             string
		inviteId         string
		expectedResponse invitestore.InviteJSON
		expectedStatus   int
	}

	responderMock := responder.NewResponder()
	invitestoreMock := invitestore.NewInviteStoreMock()
	loggerMock, _ := zap.NewProduction()
	qrMgrMock := invites_fetch.NewQRMgrMock()
	defer loggerMock.Sync()

	testcases := []testcase{
		{
			desc:     "HAPPY retrieved note correctly",
			inviteId: "1234",
			expectedResponse: invitestore.InviteJSON{
				Invite: invitestore.Invite{
					Id:    "1234",
					Title: "Hello",
				},
				QRCode: qrcodes.QRCode{},
			},
			expectedStatus: 200,
		},
		{
			desc:           "NEGATIVE not found",
			inviteId:       "12345",
			expectedStatus: 404,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			newUserMgrMock := NewInviteMgr(
				rMock,
				"dev",
				loggerMock,
				responderMock,
				invitestoreMock,
				qrMgrMock,
			)
			newUserMgrMock.Register()

			w := httptest.NewRecorder()

			r, _ := http.NewRequest("GET", BASE+"/"+testCase.inviteId, nil)
			newUserMgrMock.Router.ServeHTTP(w, r)

			var response invitestore.InviteJSON
			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
			assert.DeepEqual(t, response, testCase.expectedResponse)
		})
	}
}
