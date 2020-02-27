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
	"opman/util"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all remote Openshift objects in a Table.",
	Long: `Prints a Table containing all the requested objects found on remote server
filtered by type or labelSelector:

Syntax:
  list <flags>

Examples
  project list
  project list --types=Pod
  project list --types=Pod,ConfigMap,Service,Route,DeploymentConfig
  project list --limit=10

Flags:
  --types	Object Types to list(comma separated), e.g.: Pods,ConfigMaps,,,
  --label	Openshift Label-Selector.
  --limit	Limit number of items.
`,
	Args: cobra.RangeArgs(0, 3),
	Run: func(cmd *cobra.Command, args []string) {
		types, _ := cmd.Flags().GetString("types")
		labelSelector, _ := cmd.Flags().GetString("label")
		limit, _ := cmd.Flags().GetInt64("limit")

		err := util.ListServerItems("", Namespace, types, labelSelector, limit)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().String("types", "", "Object types like pod. Comma separated list.")
	listCmd.Flags().String("label", "", "additional Openshift Label-Selector")
	listCmd.Flags().Int64("limit", 0, "additional limitation of items returned")
}
