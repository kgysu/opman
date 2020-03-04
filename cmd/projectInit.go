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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize local sample project",
	Long: `Creates the sample project with the following files:

  - items/project.json
  - items/deployments/sample-deployment-config.json
  - items/services/sample-service.json
  - items/routes/sample-route.json
  - ...

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		folder, _ := cmd.Flags().GetString("folder")
		fmt.Printf("Init sample in './%s' \n", folder)
		err := project.InitSampleProject(folder)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println()
	},
}

func init() {
	projectCmd.AddCommand(initCmd)
	initCmd.Flags().String("folder", "items", "Path to items folder.")
}
