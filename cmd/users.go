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
	Long:  `Register with username, password, email, tel`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		tel, _ := cmd.Flags().GetString("tel")
		fmt.Println("register called with username: " + username + ", password: " + password + ", email: " + email + ", tel: " + tel)
		entity.CreateUser(username, password, email, tel)
	},
}

func init() {

	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "NULL", "Your Agenda account's username")
	registerCmd.Flags().StringP("password", "p", "NULL", "Your Agenda account's password")
	registerCmd.Flags().StringP("email", "e", "NULL", "Your Agenda account's email")
	registerCmd.Flags().StringP("tel", "t", "NULL", "Your Agenda account's telnumber")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
