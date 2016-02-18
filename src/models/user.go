package models

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var (
	uuidCount int64 = 0
	users     map[string]User
)

type User struct {
	uuid     int64
	Username string
	passHash []byte
}

func (u *User) SetPassword(newPass string) error {
	if hash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		u.passHash = hash
		return nil
	}
}

func (u *User) ValidatePassword(password string) (bool, error) {
	fmt.Println("Trying to validate ", u.passHash, "with", password)
	if err := bcrypt.CompareHashAndPassword(u.passHash, []byte(password)); err != nil {
		return false, err
	} else {
		return true, nil
	}
	return false, nil
}

func GetUserByName(name string) (User, error) {
	if _, ok := users[name]; ok {
		return users[name], nil
	} else {
		return User{}, nil
	}
}

func NewUser(un, pw string) (User, error) {
	newUser := User{uuid: uuidCount, Username: un}
	uuidCount += 1
	if err := newUser.SetPassword(pw); err != nil {
		return newUser, err
	} else {
		users[un] = newUser
		return newUser, nil
	}
}

func init() {
	users = make(map[string]User)
	NewUser("test", "password")
	NewUser("bob", "password")
	NewUser("backdoor", "password")
}
