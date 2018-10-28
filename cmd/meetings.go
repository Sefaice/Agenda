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

	"github.com/spf13/cobra"
)

// create meeting command
var cmCmd = &cobra.Command{
	Use:   "cm -t [title] -p [p1 p2 ....] -s [sTime] -e [eTime]",
	Short: "Create Meeting command",
	Long:  `Create Meeting with title, participators, start time, end time`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetString("participators")
		sTime, _ := cmd.Flags().GetString("sTime")
		eTime, _ := cmd.Flags().GetString("eTime")
		fmt.Println("meetings called with: " + title + " " + participators + " " + sTime + " " + eTime)
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)
	cmCmd.Flags().StringP("title", "t", "NULL", "New meeting's title")
	cmCmd.Flags().StringP("participators", "p", "NULL", "New meeting's participators")
	cmCmd.Flags().StringP("sTime", "s", "NULL", "New meeting's start time")
	cmCmd.Flags().StringP("eTime", "e", "NULL", "New meeting's end time")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// meetingsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// meetingsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
