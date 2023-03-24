package store

import (
	"fmt"
	"testing"
)

var (
	user_data = []struct {
		Name  string
		Email string
	}{
		{Email: "joe@bloggs.com", Name: "Joe Bloggs"},
		{Email: "jane@bloggs.com", Name: "Jane Bloggs"},
	}
)

func TestRegisterUser(t *testing.T) {

	for _, u := range user_data {
		usr, err := RegisterUser(u.Email, u.Name)

		if err != nil {
			t.Error(fmt.Sprintf("Error is not valid. Want nil, GOT: %v", err))
		}

		if usr.Name != u.Name {
			t.Error(fmt.Sprintf("User's name is not valid. Want %s, GOT: %s", u.Name, usr.Name))
		}

		if usr.Email != u.Email {
			t.Error(fmt.Sprintf("User's email is not valid. Want %s, GOT: %s", u.Email, usr.Email))
		}
	}

	// Register an existing user
	usr, err := RegisterUser(user_data[0].Email, user_data[0].Name)
	if err != UserAlreadyRegisteredErr {
		t.Error(fmt.Sprintf("Error is not valid. Want UserAlreadyRegisteredErr, GOT: %v", err))
	}
	if usr != nil {
		t.Error("User should be nil")
	}
}

func TestGetUser(t *testing.T) {
	// pre-populate the store with users
	for _, u := range user_data {
		RegisterUser(u.Email, u.Name)
	}

	// get a valid user
	usr, err := GetUser(user_data[0].Email)
	if err != nil {
		t.Error("Error should be nil")
	}

	if usr.Email != user_data[0].Email {
		t.Error(fmt.Sprintf("Invalid returned user. Want: %s, GOT: %s", user_data[0].Email, usr.Email))
	}

	// get an invalid user
	usr, err = GetUser("invalid@user.com")
	if err != UserDoesNotExistErr {
		t.Error(fmt.Sprintf("Error is not valid. Want UserDoesNotExistErr, GOT: %v", err))
	}
	if usr != nil {
		t.Error("User should be nil")
	}
}
