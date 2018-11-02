package entity

/* user.go is actually class User */

type User struct {
	Username string
	Password string
	Email    string
	Tel      string
}

func (u User) getUsername() string {
	return u.Username
}

func (u User) getPassword() string {
	return u.Password
}

func (u User) getEmail() string {
	return u.Email
}

func (u User) getTel() string {
	return u.Tel
}

func (src User) copy(tar *User) {
	tar.Username = src.Username
	tar.Password = src.Password
	tar.Email = src.Email
	tar.Tel = src.Tel
}
