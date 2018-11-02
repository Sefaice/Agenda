package entity

import (
	"io"
	"log"
	"os"
	"io/ioutil"
)

var (
	Info * log.Logger

	Warning *log.Logger

	Error * log.Logger

	Login * log.Logger
)

var errlog *os.File
var loginlog *os.File

func set(
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	loginHandle io.Writer) {
	
	Info = log.New(infoHandle,"INFO: ",log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Login = log.New(loginHandle,
		"LOG: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func init(){
	errlog = getErrLogFile()
	loginlog = gteLoginFile()

	set(ioutil.Discard,os.Stdout,errlog,loginlog)

	Error.Println("Something has failed")
	Login.Println("There is the log message")
}

func free(){
	errlog.Close()
	loginlog.Close()
}

func gteLoginFile() *os.File{
	logpath :="data/login.log"
	file,err :=os.OpenFile(logpath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	if err !=nil{
		log.Fatalf("file open error: %v ",err)
	}
	return file
}

func getErrLogFile() *os.File{
	logpath := "data/err.log"
	file, err:= os.OpenFile(logpath,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	return file
}