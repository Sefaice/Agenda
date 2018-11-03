package entity

/* service.go operates on class user*/

import (
	"fmt"
	"strings"
)

var users []User
var meetings []Meeting
var currentUser User

func init() {
	currentUser = User{ReadCurUserFromFile(), "", "", ""}
	users = UserReadFromFile()
	/*[]User{currentUser,
	User{"li", "123", "email", "tel"},
	User{"zhang", "123", "email", "tel"}}
	*/
	//_, t1 := string2ValidDate("2018-12-22-12-00")
	//_, t2 := string2ValidDate("2018-12-23-12-00")
	meetings = MeetingReadFromFile()
	//[]Meeting{Meeting{"m1", "wang", []string{"li", "zhang"}, t1, t2}}
}

/**
 * users operation
 */

//func name MUST start with an upper case if you want to call it in other packages
func CreateUser(username string, password string, email string, tel string) {
	if username == "" || password == "" || email == "" || tel == "" {
		fmt.Println("Paramaters can't be empty!")
		return
	}
	for i := 0; i < len(users); i++ {
		if users[i].getUsername() == username {
			fmt.Println("Username: " + username + " is existed!")
			Error.Println("Username: " + username + " is existed!")
			return
		}
	}
	u := User{username, password, email, tel}
	users = append(users, u)
	//fmt.Println("usersasdasdasdasd", users)
	UserWriteFile(users)
	fmt.Println("Create user: " + username + " success! ")
	Login.Println("Create user: " + username + " success! ")
}

func UserLogin(username string, password string) {
	//fmt.Println(currentUser.getUsername())
	if currentUser.getUsername() != "" {
		fmt.Println("you should log out first")
		Login.Println("try to log in repeatedly")
		return
	}
	if username == "" || password == "" {
		fmt.Println("Paramaters can't be empty!")
		return
	}
	for i := 0; i < len(users); i++ {
		if users[i].getUsername() == username {
			if users[i].getPassword() == password {
				users[i].copy(&currentUser)
				fmt.Println("User: " + username + " login success!")
				Login.Println("User: " + username + " login success!")
				CurUserWriteFile(username)
			} else {
				fmt.Println("Password is wrong!")
				Error.Println("Password is wrong!")
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
	Login.Println("User: " + currentUser.getUsername() + " log out seccess!")
	CurUserFileDelete()
	currentUser = User{"", "", "", ""}
}

func PrintAllUsers() {
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i].getUsername() + "   " + users[i].getEmail() + "    " + users[i].getTel())
		Login.Println(currentUser.getUsername() + "print all users")
	}
}

func DeleteUser() {
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	quitAllMeetings()
	pos := indexOfUsers()
	if pos < 0 {
		fmt.Println("User: " + currentUser.getUsername() + " is not exist!")
		return
	}
	users = append(users[:pos], users[pos+1:]...)
	fmt.Println("User: " + currentUser.getUsername() + " delete seccess!")
	Login.Println("delete user" + currentUser.getUsername())
	currentUser = User{"", "", "", ""}
	CurUserFileDelete()
	UserWriteFile(users)
}

/**
 * users bridge funcs
 */

func indexOfUsers() int {
	for i := 0; i < len(users); i++ {
		if users[i].getUsername() == currentUser.getUsername() {
			return i
		}
	}
	return -1
}

func getIndexOfUserByName(username string) int {
	for i := 0; i < len(users); i++ {
		if users[i].getUsername() == username {
			return i
		}
	}
	return -1
}

/**
 * meetings operation
 */

func CreateMeeting(title string, pStr string, sStr string, eStr string) {
	if title == "" || pStr == "" || sStr == "" || eStr == "" {
		fmt.Println("Paramaters can't be empty!")
		return
	}
	//need login
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	//title distinct
	for i := 0; i < len(meetings); i++ {
		if meetings[i].getTitle() == title {
			fmt.Println("Title: " + title + " is existed!")
			return
		}
	}
	//participators must be Agenda users
	pValid, pArr := parseParticipators(pStr)
	if !pValid {
		return
	}
	//sTime eTime valid
	sValid, sDate := string2ValidDate(sStr)
	if !sValid {
		fmt.Println("Start time " + sStr + " invalid!")
		return
	}
	eValid, eDate := string2ValidDate(eStr)
	if !eValid {
		fmt.Println("End time " + eStr + " invalid!")
		return
	}
	if !compareDate(sDate, eDate) {
		fmt.Println("Start time should be earlier than end time!")
		return
	}
	//sponsor time available
	if !timeAvailable(currentUser, sDate, eDate) {
		fmt.Println("Sponsor time not available!")
		return
	}
	//participators' time available
	if !pTimeAvailable(pArr, sDate, eDate) {
		fmt.Println("Participators time not available!")
		return
	}
	//success
	m := Meeting{title, currentUser.getUsername(), pArr, sDate, eDate}
	meetings = append(meetings, m)
	fmt.Println("Create meeting: " + title + " success! ")
	Login.Println("creating meeting" + title + "successfully!")
	MeetingWriteFile(meetings)
}

func AddParticipators(title string, pStr string) {
	if title == "" || pStr == "" {
		fmt.Println("Paramaters can't be empty!")
		return
	}
	//need login
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	//meeting exists and must be sponsor
	index := getIndexOfMeetingByTitle(title)
	if index == -1 {
		fmt.Println("Meeting " + title + " does not exist!")
		return
	}
	//use index instead
	if meetings[index].getSponsor() != currentUser.getUsername() {
		fmt.Println("You are not sponsor of this meeting!")
		return
	}
	//participators must be Agenda users
	pValid, pArr := parseParticipators(pStr)
	if !pValid {
		return
	}
	//not already sponsor or participator
	for i := 0; i < len(pArr); i++ {
		if meetings[index].isParticipatorOrSponsor(users[getIndexOfUserByName(pArr[i])]) {
			fmt.Println(pArr[i] + " is already a sponsor or participator of this meeting!")
			return
		}
	}
	//time available
	if !pTimeAvailable(pArr, meetings[index].getStart(), meetings[index].getEnd()) {
		fmt.Println("Participators time not available!")
		return
	}
	//success
	meetings[index].addParticipators(pArr)
	fmt.Println("Add participators " + getParticipatorsStr(pArr) + " success!")
	Login.Println(title + "add participator " + pStr + " successfully")
	MeetingWriteFile(meetings)
}

func DeleteParticipators(title string, pStr string) {
	if title == "" || pStr == "" {
		fmt.Println("Paramaters can't be empty!")
		return
	}
	//need login
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	//meeting exists and must be sponsor
	index := getIndexOfMeetingByTitle(title)
	if index == -1 {
		fmt.Println("Meeting " + title + " does not exist!")
		return
	}
	if meetings[index].getSponsor() != currentUser.getUsername() {
		fmt.Println("You are not sponsor of this meeting!")
		return
	}
	//participators must be Agenda users
	pValid, pArr := parseParticipators(pStr)
	if !pValid {
		return
	}
	//cannot be sponsor
	for i := 0; i < len(pArr); i++ {
		if meetings[index].getSponsor() == pArr[i] {
			fmt.Println(pArr[i] + " is already a sponsor of this meeting!")
			return
		}
	}
	//must already a participator
	for i := 0; i < len(pArr); i++ {
		if !meetings[index].isParticipatorOrSponsor(users[getIndexOfUserByName(pArr[i])]) {
			fmt.Println(pArr[i] + " is not a participator of this meeting!")
			return
		}
	}
	//success
	meetings[index].deleteParticipators(pArr)
	fmt.Println("Delete participators " + getParticipatorsStr(pArr) + " success!")
	Login.Println(title + "delete participator " + pStr + " successfully ")
	MeetingWriteFile(meetings)
}

func QueryMeetings(sStr string, eStr string) {
	if sStr == "" || eStr == "" {
		fmt.Println("Paramaters can't be empty!")
		return
	}
	//need login
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	//time valid
	sValid, sDate := string2ValidDate(sStr)
	if !sValid {
		fmt.Println("Start time " + sStr + " invalid!")
		return
	}
	eValid, eDate := string2ValidDate(eStr)
	if !eValid {
		fmt.Println("End time " + eStr + " invalid!")
		return
	}
	if !compareDate(sDate, eDate) {
		fmt.Println("Start time should be earlier than end time!")
		return
	}
	//query
	for _, m := range meetings {
		if m.isParticipatorOrSponsor(currentUser) {
			ms := m.getStart()
			me := m.getEnd()
			// (ms <= s < me) || (ms < e <= me) || (s <= ms < e) || (s < me <=e)
			if (!compareDate(sDate, ms) && compareDate(sDate, me)) ||
				(compareDate(ms, eDate) && (compareDate(eDate, me) || equalDate(eDate, me))) ||
				(!compareDate(ms, sDate) && compareDate(ms, eDate)) ||
				(compareDate(sDate, me) && (compareDate(me, eDate) || equalDate(me, eDate))) {
				m.printMeeting()
			}
		}
	}
}

func DeleteMeeting(title string) {
	if title == "" {
		fmt.Println("Paramater can't be empty!")
		return
	}
	//need login
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	//meeting exists and must be sponsor
	index := getIndexOfMeetingByTitle(title)
	if index == -1 {
		fmt.Println("Meeting " + title + " does not exist!")
		return
	}
	m := meetings[index]
	if m.getSponsor() != currentUser.getUsername() {
		fmt.Println("You are not sponsor of this meeting!")
		return
	}
	//success
	meetings = append(meetings[:index], meetings[:index+1]...)
	fmt.Println("Delete meeting " + title + " success!")
	Login.Println("delete meeting " + title + " successfully")
	MeetingWriteFile(meetings)
}

func QuitMeeting(title string) {
	if title == "" {
		fmt.Println("Paramater can't be empty!")
		return
	}
	//need login
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	//meeting exists
	index := getIndexOfMeetingByTitle(title)
	if index == -1 {
		fmt.Println("Meeting " + title + " does not exist!")
		return
	}
	m := meetings[index]
	//must be sponsor or participator
	if !m.isParticipatorOrSponsor(currentUser) {
		fmt.Println("You are not sponsor or participator of thie meeting!")
		return
	}
	//quit the meeting
	if m.getSponsor() == currentUser.getUsername() {
		meetings = append(meetings[:index], meetings[:index+1]...)
		fmt.Println("Quit meeting " + title + " success and the meeting is deleted!")
	} else {
		m.deleteParticipators([]string{currentUser.getUsername()})
		fmt.Println("Quit meeting " + title + " success!")
	}

	MeetingWriteFile(meetings)
}

func ClearAllMeetings() {
	//need login
	if len(currentUser.getUsername()) <= 0 {
		fmt.Println("Login First!")
		return
	}
	MeetingWriteFile(meetings)
}

/**
 * meetings bridge funcs
 */

func getIndexOfMeetingByTitle(title string) int {
	for i := 0; i < len(meetings); i++ {
		if meetings[i].getTitle() == title {
			return i
		}
	}
	return -1
}

func parseParticipators(pStr string) (bool, []string) {
	pArr := strings.Split(pStr, ",")
	for i := 0; i < len(pArr); i++ {
		if getIndexOfUserByName(pArr[i]) == -1 {
			fmt.Println("Participator " + pArr[i] + " is not Agenda user!")
			return false, []string{}
		}
	}
	return true, pArr
}

//TAKE DATES VALID
func timeAvailable(u User, sDate Date, eDate Date) bool {
	//get u's all meetings and check each one for date
	for _, m := range meetings {
		if m.isParticipatorOrSponsor(u) {
			ms := m.getStart()
			me := m.getEnd()
			// (ms <= s < me) || (ms < e <= me) || (s <= ms < e) || (s < me <=e)
			if (!compareDate(sDate, ms) && compareDate(sDate, me)) ||
				(compareDate(ms, eDate) && (compareDate(eDate, me) || equalDate(eDate, me))) ||
				(!compareDate(ms, sDate) && compareDate(ms, eDate)) ||
				(compareDate(sDate, me) && (compareDate(me, eDate) || equalDate(me, eDate))) {
				return false
			}
		}
	}
	return true
}

//TAKE PARR, SDATE, EDATE AS VALID
func pTimeAvailable(pArr []string, sDate Date, eDate Date) bool {
	for i := 0; i < len(pArr); i++ {
		index := getIndexOfUserByName(pArr[i])
		if index != -1 {
			if !timeAvailable(users[index], sDate, eDate) {
				return false
			}
		}
	}
	return true
}

//quit all meetings for current user
func quitAllMeetings() {
	for _, m := range meetings {
		if m.getSponsor() == currentUser.getUsername() {
			index := getIndexOfMeetingByTitle(m.getTitle())
			meetings = append(meetings[:index], meetings[:index+1]...)
			fmt.Println("Quit meeting " + m.getTitle() + " success and the meeting is deleted!")
		} else {
			m.deleteParticipators([]string{currentUser.getUsername()})
			fmt.Println("Quit meeting " + m.getTitle() + " success!")
		}
	}

	MeetingWriteFile(meetings)
}
