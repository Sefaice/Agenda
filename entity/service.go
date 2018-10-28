package entity

/* service.go operates on class user*/

import "fmt"

//func name MUST start with an upper case if you want to call it in other packages
func CreateUser(username string, password string, email string, tel string) {

	u := User{username, password, email, tel}

	fmt.Println("in service: " + u.getUsername())
}
