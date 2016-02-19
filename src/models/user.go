package models

// Basic User Interface

type User interface {
	SetPassword(string) error
	ValidatePassword(string) (bool, error)
	Name() string
}

type UserRepository interface {
	GetUserByName(name string) (User, error)
	NewUser(un, pw string) (User, error)
}
