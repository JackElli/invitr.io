package userstore

type UserStoreMock struct{}

func NewUserStoreMock() *UserStoreMock {
	return &UserStoreMock{}
}

func (store *UserStoreMock) Get(id string) (*User, error) {
	return &User{
		Id:        "1234",
		FirstName: "Jack",
		LastName:  "Ellis",
	}, nil
}

func (store *UserStoreMock) Insert(user *User) (*User, error) {
	return user, nil
}

func (store *UserStoreMock) Remove(userId string) error {
	return nil
}
