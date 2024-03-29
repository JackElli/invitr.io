package user

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
	"invitr.io.com/services/users/userstore"
)

func TestNewUser(t *testing.T) {
	const ROOT_TEST = "/user"

	type testcase struct {
		desc             string
		expectedResponse userstore.User
		expectedStatus   int
	}

	responderMock := responder.NewResponder()
	userstoreMock := userstore.NewUserStoreMock()
	loggerMock, _ := zap.NewProduction()
	defer loggerMock.Sync()

	mockId := "1234"

	newUser := userstore.User{
		Id:        &mockId,
		FirstName: "Jack",
		LastName:  "Ellis",
	}

	testcases := []testcase{
		{
			desc: "HAPPY added user correctly",
			expectedResponse: userstore.User{
				Id:        &mockId,
				FirstName: "Jack",
				LastName:  "Ellis",
			},
			expectedStatus: 201,
		},
	}

	for _, testCase := range testcases {
		t.Run(testCase.desc, func(t *testing.T) {
			rMock := mux.NewRouter()

			newUserMgrMock := NewUserMgr(
				rMock,
				loggerMock,
				responderMock,
				userstoreMock,
			)
			newUserMgrMock.Register()

			w := httptest.NewRecorder()

			newUserData, _ := json.Marshal(newUser)

			r, _ := http.NewRequest("POST", ROOT_TEST, bytes.NewBuffer(newUserData))
			newUserMgrMock.Router.ServeHTTP(w, r)

			var response userstore.User
			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, w.Result().StatusCode, testCase.expectedStatus)
			assert.DeepEqual(t, response, testCase.expectedResponse)
		})
	}
}
