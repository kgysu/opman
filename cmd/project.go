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
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "work with your local project (default './items')",
	Long: `Manage your local project in './items':

First initialize a new local project. Then you can list all found objects and verify all is setup correctly.
Now you are ready to deploy/remove/start/stop/compare these objects.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("specify command")
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
