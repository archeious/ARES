package user

import "fmt"

type BaseUserRepository struct {
	uuidCount int64
	users     map[string]BaseUser
}

func (b *BaseUserRepository) GetUserByName(name string) (User, error) {
	if _, ok := b.users[name]; ok {
		u := b.users[name]
		return &u, nil
	} else {
		return &BaseUser{}, nil
	}
}

func (b *BaseUserRepository) NewUser(un, pw string) (User, error) {
	newUser := BaseUser{uuid: uuidCount, Username: un}
	uuidCount += 1
	if err := newUser.SetPassword(pw); err != nil {
		return &newUser, err
	} else {
		b.users[un] = newUser
		fmt.Printf("\n\n\n %v \n\n\n", *b)
		return &newUser, nil
	}
}

func NewBaseUserRepository() *BaseUserRepository {
	bur := BaseUserRepository{uuidCount: 0, users: make(map[string]BaseUser)}
	return &bur
}
