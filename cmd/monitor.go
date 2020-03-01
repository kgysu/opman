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
	"opman/util"
)

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor openshift objects and show diffs in real time.",
	Long: `
Monitor:
`,
	Run: func(cmd *cobra.Command, args []string) {
		types, _ := cmd.Flags().GetString("types")
		labelSelector, _ := cmd.Flags().GetString("label")
		limit, _ := cmd.Flags().GetInt64("limit")

		err := util.MonitorRemoteItems("", Namespace, types, labelSelector, limit)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.Flags().String("types", "", "Object types like pod. Comma separated list.")
	monitorCmd.Flags().String("label", "", "additional Openshift Label-Selector")
	monitorCmd.Flags().Int64("limit", 0, "additional limitation of items returned")
}
