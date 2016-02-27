package user

import "fmt"

type MemUserRepository struct {
	uuidCount int64
	users     map[string]BaseUser
}

func (b *MemUserRepository) GetUserByName(name string) (User, error) {
	if _, ok := b.users[name]; ok {
		fmt.Println("\nGETTING: GetUserByName with %v from %v", name, b)
		u := b.users[name]
		return &u, nil
	} else {
		return &BaseUser{}, nil
	}
}

func (b *MemUserRepository) NewUser(un, pw string) (User, error) {
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

func NewMemUserRepository() *BaseUserRepository {
	bur := BaseUserRepository{uuidCount: 0, users: make(map[string]BaseUser)}
	return &bur
}
