package behavior_pattern

import "fmt"

type User struct {
	name  string
	phone *Phone
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) Dail(word string) {
	fmt.Print("User -> ", u.name)
	u.phone.Tell(word)
}

type Phone struct {
	user *User
}

func NewPhone() *Phone {
	return &Phone{}
}

func (p *Phone) Tell(word string) {
	fmt.Println(" tell -> ", word)
}
