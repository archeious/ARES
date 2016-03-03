package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var (
	uuidCount int64 = 0
	users     map[string]BaseUser
)

type BaseUser struct {
	uuid     int64
	Username string
	passHash []byte
}

func (u *BaseUser) SetPassword(newPass string) error {
	if hash, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		u.passHash = hash
		return nil
	}
}

func (u *BaseUser) ValidatePassword(password string) (bool, error) {
	fmt.Println("Trying to validate ", u.passHash, "with", password)
	if err := bcrypt.CompareHashAndPassword(u.passHash, []byte(password)); err != nil {
		return false, err
	}
	return true, nil

}

func (u *BaseUser) Name() string {
	return u.Username
}
