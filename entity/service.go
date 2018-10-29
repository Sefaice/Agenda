package entity

/* service.go operates on class user*/

import "fmt"

var users []User
var meetings []Meeting
var currentUser User

func init() {
	currentUser = User{"", "", "", ""}
}

//func name MUST start with an upper case if you want to call it in other packages
func CreateUser(username string, password string, email string, tel string) {
	if username == "" || password == "" || email == "" || tel == "" {
		fmt.Println("Paramater can't be empty!")
		return
	}
	for i := 0; i < len(users); i++ {
		if users[i].getUsername() == username {
			fmt.Println("Username: " + username + " is existed!")
		}
	}
	u := User{username, password, email, tel}
	users = append(users, u)
	fmt.Println("Create user: " + username + " success! ")
}

func UserLogin(username string, password string) {
	for i := 0; i < len(users); i++ {
		if users[i].getUsername() == username {
			if users[i].getPassword() == password {
				users[i].copy(&currentUser)
				fmt.Println("User: " + username + " login success!")
			} else {
				fmt.Println("Password is wrong!")
			}
			return
		}
	}
	fmt.Println("User: " + username + " is not exist!")
}

func UserLogout() {
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	fmt.Println("User: " + currentUser.getUsername() + " log out seccess!")
	currentUser = User{"", "", "", ""}
}

func PrintAllUsers() {
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i].getUsername() + "   " + users[i].getEmail() + "    " + users[i].getTel())
	}
}

func DeleteUser() {
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	QuitAllMeetings()
	pos := IndexOfUsers()
	if pos < 0 {
		fmt.Println("User: " + currentUser.getUsername() + " is not exist!")
		return
	}
	users = append(users[:pos], users[pos+1:]...)
	fmt.Println("User: " + currentUser.getUsername() + " delete seccess!")
	currentUser = User{"", "", "", ""}
}

func IndexOfUsers() int {
	for i := 0; i < len(users); i++ {
		if users[i].getUsername() == currentUser.getUsername() {
			return i
		}
	}
	return -1
}
