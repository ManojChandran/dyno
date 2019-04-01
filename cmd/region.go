// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"github.com/spf13/cobra"
)

// regionCmd represents the region command
var regionCmd = &cobra.Command{
	Use:   "region",
	Short: "Set AWS region",
	Long: `Set the AWS to access the Table
For example:
dyno region <table name>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("region required,\n refer https://docs.aws.amazon.com/general/latest/gr/rande.html")
			os.Exit(0)
		}
		os.Setenv("AWS_REGION", args[0])
		fmt.Println("AWS region set to ", os.Getenv("AWS_REGION"))
	},
}

func init() {
	rootCmd.AddCommand(regionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// regionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// regionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
