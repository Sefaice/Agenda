package entity

/* user.go is actually class User */

type User struct {
	username string
	password string
	email    string
	tel      string
}

func (u User) getUsername() string {
	return u.username
}

func (u User) getPassword() string {
	return u.password
}

func (u User) getEmail() string {
	return u.email
}

func (u User) getTel() string {
	return u.tel
}

func (src User) copy(tar *User) {
	tar.username = src.username
	tar.password = src.password
	tar.email = src.email
	tar.tel = src.tel
}
