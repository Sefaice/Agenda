# GO-Agenda

基于go语言cobra包的命令行会议管理系统，服务计算课程项目。

## How to run

要求配置了go语言以及cobra包，cobra包在下载时可能遇到网络问题，具体步骤如下：

>使用命令 `go get -v github.com/spf13/cobra/cobra` 下载过程中会报错，请在 $GOPATH/src/golang.org/x 目录下用 `git clone https://github.com/go-lang/sys` 和 `git clone https://github.com/go-lang/test`下载sys和test项目，然后使用`go install github.com/spf13/cobra/cobra`安装即可。

环境配置完毕，在项目目录下使用命令`go run main.go [command] [flags] [options]` 即可运行项目，比如：

`go run main.go register -u fangege -p 123456 -e fanruichao1@live.com -t 15866668888'

就在系统中注册了一个新用户，接下来用`go run main.go login -u fangege -p 123456`即可登陆系统。

## 项目结构

文件结构如下：

```
├── cmd
│   ├── meetings.go
│   ├── root.go
│   └── users.go
├── cmd-design.md
├── entity
│   ├── Date.go
│   ├── Meeting.go
│   ├── service.go
│   ├── Storage.go
│   └── user.go
├── LICENSE
├── main.go
└── README.md
```

cmd-design.md是项目的命令设计；main.go是整个项目的入口文件；/cmd下存放的是用cobra处理命令输入的go文件；/entity下是处理逻辑操作的go文件，其中meeting.go,user.go,service.go是用方法实现的go语言类文件，storage.go用于操作文件读写，service.go是逻辑操作的入口，提供方法供cmd调用。

项目采用的是MVC结构，cmd下是view视图层，entity中基本类文件是model层，service.go充当了controller层，结构很简单，也很清晰，所以容易分工且耦合简单。

## 基础知识

### go命令行

本项目用到了cobra包处理命令行输入，上一次作业使用的pflag包使用是对go命令行规范的熟悉，cobra的处理方式类似，只需要熟悉它的文件结构即可。

### go面向对象

[老师的面向对象讲解](https://pmlpml.github.io/ServiceComputingOnCloud/oo-thinking#3%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F%E5%BA%94%E7%94%A8---command-%E5%AF%B9%E8%B1%A1%E8%AE%BE%E8%AE%A1)

go语言并没有class类型，但是可以通过为结构体添加方法进行实现，比如user.go的设计：

```
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
```

其中的func其实是go语言中的方法method，通过receiver argument实现给类添加方法，这样就构成了包含数据和方法的user类，meeting和date同理

## 实现过程

首先进行结构设计，确定采用MVC模式，确定文件结构。

然后根据业务需求填写[cmd-design.md](https://github.com/Sefaice/Agenda/blob/master/cmd-design.md)中的命令和参数设计，明确每条指令的运行条件和参数要求。

实现按照View-Model-Controller的顺序进行。

### View

/cmd文件下使用cobra的添加指令的命令加入meetings.go和users.go两个文件，分别处理会议和用户操作的命令输入，因为这里只需要负责读取参数和传给控制层，所以每条命令的实现模式完全一样，不需要进行验证：
```
// register command
var registerCmd = &cobra.Command{
	//register -u(--username) [yourUsername] -p(--password) [yourPassword] -e(--email) [yourEmail] -t(--tel) [youTelnumber]
	Use:   "register -u [username] -p [password] -e [email] -t [tel]",
	Short: "Register Command",
	Long:  "Register with username, password, email, tel",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		tel, _ := cmd.Flags().GetString("tel")
		fmt.Println("register called with username: " + username + ", 			password: " + password + ", email: " + email + ", tel: " + tel)
		entity.CreateUser(username, password, email, tel)
	},
}
```

然后在文件中的init（）方法添加对应命令即可：
```
rootCmd.AddCommand(registerCmd)
registerCmd.Flags().StringP("username", "u", "", "Your Agenda account's username")
registerCmd.Flags().StringP("password", "p", "", "Your Agenda account's password")
registerCmd.Flags().StringP("email", "e", "", "Your Agenda account's email")
registerCmd.Flags().StringP("tel", "t", "", "Your Agenda account's telnumber")
```

### Model

Model层包含三个“类文件”，User.go， Meeting.go，Date.go，采用上述结构体方法的办法实现，添加对应单个类的操作方法，比如修改一个会议的与会者。
```
//TAKE PARR ALL AS VALID USERS
func (m Meeting) addParticipators(pArr []string) {
	for _, p := range pArr {
		m.participators = append(m.participators, p)
	}
}

//TAKE PARR ALL AS VALID USERS
func (m Meeting) deleteParticipators(pArr []string) {
	for _, p := range pArr {
		for i, q := range m.participators {
			if q == p {
				m.participators = append(m.participators[:i], m.participators[i+1:]...)
			}
		}
	}

	fmt.Println(m.participators)
}

func getParticipatorsStr(pArr []string) string {
	pStr := ""
	for _, p := range pArr {
		if pStr == "" {
			pStr = p
		} else {
			pStr = pStr + ", " + p
		}
	}
	return pStr
}
```

### Controller

控制层处在视图层和模型层中间，为每一条cmd指令提供函数，实现在Service.go中，其中每条命令的实现步骤基本类似：

* 验证参数格式，不能为空
* 验证是否登陆（如果需要）
* 解析参数，验证参数合法性
* 进行命令操作，写入文件，写入log包

比如为某个会议添加与会者的实现：
```

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
	m := meetings[index]
	if m.getSponsor() != currentUser.getUsername() {
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
		if m.isParticipatorOrSponsor(users[getIndexOfUserByName(pArr[i])]) {
			fmt.Println(pArr[i] + " is already a sponsor or participator of this meeting!")
			return
		}
	}
	//time available
	if !pTimeAvailable(pArr, m.getStart(), m.getEnd()) {
		fmt.Println("Participators time not available!")
		return
	}
	//success
	m.addParticipators(pArr)
	fmt.Println("Add participators " + getParticipatorsStr(pArr) + " success!")

	MeetingWriteFile(meetings);
}
```

至此，整个项目的实现就完成了，接下来对项目进行测试和修改即可。
