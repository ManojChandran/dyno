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
	"github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Purge the dynamo table",
	Long: `Purge dynamodb table
For example:
dyno purge <table name>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			 fmt.Println("Table Name missing")
			 os.Exit(0)
		}

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION"))},
		)
		// Create DynamoDB client
    svc := dynamodb.New(sess)
		input := &dynamodb.DeleteTableInput{
			TableName: aws.String(args[0]),
		}

    // Get the list of tables
		_, err = svc.DeleteTable(input)
		if err != nil {
				fmt.Println("Got error calling DeleteTable:")
				fmt.Println(err.Error())
				os.Exit(1)
		}
		fmt.Println("requested table deleted from ",os.Getenv("AWS_REGION"))
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// purgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// purgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
