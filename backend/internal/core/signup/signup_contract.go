package signup

type SignupContract interface {
	Signup(email string, password string, name string) error
}

type SignupRepositoryContract interface {
	StoreUser(email string, password string, name string, salt string, userType string, score float32) error
}