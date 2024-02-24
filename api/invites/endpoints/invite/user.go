package invite

import (
	"encoding/json"
	"net/http"

	"invitio.com/users/userstore"
)

const (
	USER_URL = "http://users:3200/user/"
)

func GetUser(userId string) (*userstore.User, error) {
	resp, err := http.Get(USER_URL + userId)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user userstore.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
