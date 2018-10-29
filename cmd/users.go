// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/sefaice/Agenda/entity"
	"github.com/spf13/cobra"
)

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
		fmt.Println("register called with username: " + username + ", password: " + password + ", email: " + email + ", tel: " + tel)
		entity.CreateUser(username, password, email, tel)
	},
}

// login command
var loginCmd = &cobra.Command{
	Use:   "login -u [username] -p [password]",
	Short: "Login Command",
	Long:  "Login with username and password",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		fmt.Println("login called with username: " + username + ", password: " + password)
		entity.UserLogin(username, password)
	},
}

// logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout Command",
	Long:  "Logout account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logout called")
		entity.UserLogout()
	},
}

// list users command
var luCmd = &cobra.Command{
	Use:   "lu",
	Short: "List All Users Command",
	Long:  "List all users if you already login",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lu called")
		entity.PrintAllUsers()
	},
}

// delete account command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Account Command",
	Long:  "Delete your account if you already login",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		entity.DeleteUser()
	},
}

func init() {

	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "Your Agenda account's username")
	registerCmd.Flags().StringP("password", "p", "", "Your Agenda account's password")
	registerCmd.Flags().StringP("email", "e", "", "Your Agenda account's email")
	registerCmd.Flags().StringP("tel", "t", "", "Your Agenda account's telnumber")

	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "Your Agenda account's username")
	loginCmd.Flags().StringP("password", "p", "", "Your Agenda account's password")

	rootCmd.AddCommand(logoutCmd)

	rootCmd.AddCommand(luCmd)

	rootCmd.AddCommand(deleteCmd)
}
