package usermgr

type UserStoreMock struct{}

func NewUserStoreMock() *UserStoreMock {
	return &UserStoreMock{}
}

func (store *UserStoreMock) Get(id string) (*User, error) {
	return &User{
		UserId:   "1234",
		Username: "Jack",
	}, nil
}

func (store *UserStoreMock) Insert(user *User) (*User, error) {
	return &User{
		UserId:   "1234",
		Username: "Jack",
	}, nil
}

func (store *UserStoreMock) Remove(userId string) error {
	return nil
}
