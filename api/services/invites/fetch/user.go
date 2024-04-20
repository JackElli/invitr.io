package invites_fetch

import (
	"encoding/json"
	"errors"
	"net/http"

	"invitr.io.com/services/users/userstore"
)

const USER_URL = "http://users:3200/user/"

// GetUser fetches a user based on an ID from the user service
func GetUser(userId string) (*userstore.User, error) {
	resp, err := http.Get(USER_URL + userId)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("user not found, cannot create invite")
	}

	var user userstore.User
	json.NewDecoder(resp.Body).Decode(&user)

	return &user, nil
}
