package entity

/* user.go is actually class User */

type User struct {
	username string
	password string
	email    string
	tel      string
}

func (usr User) getUsername() string {
	return usr.username
}
