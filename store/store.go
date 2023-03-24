package store

import "crypto_task/wallet"

var (
	_users = make(map[string]*wallet.User)
)

// RegisterUser will register a new user with the store. If user with the same email  already exists then
// UserAlreadyRegisteredErr is returned, otherwise a new user with  a wallet of zero balance is created
// and inserted into the store
func RegisterUser(email, name string) (*wallet.User, error) {
	if _, ok := _users[email]; ok {
		return nil, UserAlreadyRegisteredErr
	}

	user := wallet.NewUser(email, name)

	_users[email] = user
	return user, nil
}

// GetUser will return a user with provided email  that is stored in the store. If the user with provided email
// does not exist then UserDoesNotExistErr is returned, otherwise a valid user is returned
func GetUser(email string) (*wallet.User, error) {
	if user, ok := _users[email]; ok {
		return user, nil
	}
	return nil, UserDoesNotExistErr
}
