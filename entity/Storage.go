package entity

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var mData []Meeting
var uData []User

//when the program is first run, the user and meeting info will be read from file at first.
//so that we have a copy of info in the file in memory. and service.go can query in this file.
func Init() {
	tmpM := MeetingReadFromFile()
	tmpU := UserReadFromFile()

	for i := 0; i < len(tmpM); i++ {
		mData = append(mData, tmpM[i])
	}

	for i := 0; i < len(tmpU); i++ {
		uData = append(uData, tmpU[i])
	}
}

//validation should be executed in controller(service) and this file only concerns about read from file and write to file
//I first think the info array should be in storage.go but zwz and frc has implmented major functions in service.go
//so
func UserWriteFile(userArr []User) {
	//fmt.Println("fucksadsads", userArr)
	file, err := os.OpenFile("entity/data/User.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	os.Truncate("entity/data/User.txt", 0)
	if err != nil {
		fmt.Println("error open file")
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(userArr); i++ {
		//fmt.Println(userArr[i])
		file.WriteString(string(UserJsonEncode(userArr[i])))
		if i != len(userArr)-1 {
			file.WriteString("\n")
		}
	}

}

func UserReadFromFile() []User {
	var tmp []User
	f, err := os.Open("entity/data/User.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		//fmt.Println("user read: " + line)
		if err != nil && io.EOF == err {
			//read finish
			tmp = append(tmp, UserJsonDecode([]byte(line)))
			break
		}

		tmp = append(tmp, UserJsonDecode([]byte(line)))
	}
	return tmp
}

func CurUserWriteFile(curUser string) {
	file, err := os.OpenFile("entity/data/curUser.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		Error.Println("write curuser to file failed")
		os.Exit(1)
	}
	defer file.Close()
	//file.WriteString(string(UserJsonEncode()))
	file.WriteString(curUser)
}

func CurUserFileDelete() {
	file, err := os.OpenFile("entity/data/curUser.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("opening curuser file failed")
	}
	defer file.Close()

	file.WriteString("")
}

func ReadCurUserFromFile() string {
	//var tmp User;
	f, err := os.Open("entity/data/curUser.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	bytesLine, _, err := rd.ReadLine()
	line := string(bytesLine)
	//fmt.Println("storage read cur usr: " + line)
	if err == io.EOF {
		return ""
	}

	return line
}

func MeetingReadFromFile() []Meeting {
	var tmp []Meeting
	f, err := os.Open("entity/data/Meeting.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)

	for {
		bytesLine, _, err := rd.ReadLine()
		line := string(bytesLine)
		if err == io.EOF {
			if line != "" {
				tmp = append(tmp, MeetingJsonDecode([]byte(line)))
			} else {
				break
			}
		}
		tmp = append(tmp, MeetingJsonDecode([]byte(line)))
	}

	return tmp
}

func MeetingWriteFile(meetingArr []Meeting) {
	file, err := os.OpenFile("entity/data/Meeting.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	os.Truncate("entity/data/Meeting.txt", 0)

	if err != nil {
		fmt.Println("metting write error")
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(meetingArr); i++ {
		file.WriteString(string(MeetingJsonEncode(meetingArr[i])[:]))
		file.WriteString("\n")
	}

}

func UserJsonEncode(m User) []byte {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error user json")
		os.Exit(1)
	}
	return data
}

func UserJsonDecode(js []byte) User {
	var ju User
	err := json.Unmarshal(js, &ju)
	if err != nil {
		fmt.Println("error user decode")
	}
	return ju
}

func MeetingJsonEncode(m Meeting) []byte {
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("error meeting json")
		os.Exit(1)
	}
	return data
}

func MeetingJsonDecode(js []byte) Meeting {
	var jm Meeting
	err := json.Unmarshal(js, &jm)
	if err != nil {
		fmt.Println("error meeting decode")
	}
	return jm
}
