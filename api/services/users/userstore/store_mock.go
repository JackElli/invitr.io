package userstore

type UserStoreMock struct{}

func NewUserStoreMock() *UserStoreMock {
	return &UserStoreMock{}
}

func (store *UserStoreMock) Get(id string) (*User, error) {
	mockId := "1234"
	return &User{
		Id:        &mockId,
		FirstName: "Jack",
		LastName:  "Ellis",
	}, nil
}

func (store *UserStoreMock) Insert(user *User) (*User, error) {
	mockId := "1234"
	return &User{
		Id:        &mockId,
		FirstName: "Jack",
		LastName:  "Ellis",
	}, nil
}

func (store *UserStoreMock) Remove(userId string) error {
	return nil
}
