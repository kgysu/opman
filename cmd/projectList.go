/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"opman/util/project"
)

// listCmd represents the list command
var listProjectCmd = &cobra.Command{
	Use:   "list",
	Short: "list all files",
	Long: `List all project files found in './<folder>':

Syntax:
  project list <flags>

Examples
  project list 
  project list --folder=something
  project list -d
  project list --debug

Flags:
  --folder	path to local project
  --debug	debug mode on

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		folder, _ := cmd.Flags().GetString("folder")
		fmt.Printf("List Objects found in './%s' \n", folder)
		debug, _ := cmd.Flags().GetBool("debug")
		if debug {
			fmt.Println("Running in Debug-Mode.")
		}

		err := project.ListLocalProjectFiles(Namespace, folder, debug)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println()
	},
}

func init() {
	projectCmd.AddCommand(listProjectCmd)
	listProjectCmd.Flags().String("folder", "items", "Path to items folder.")
	listProjectCmd.Flags().BoolP("debug", "d", false, "Debug-Mode: create output file to determine result.")
}
