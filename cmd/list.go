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

	"github.com/spf13/cobra"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the tables in the DynamoDB",
	Long: `list all the tables in DynamoDB.
	For example:

list will give all the table names currently available in the AWS account.
It list out the table name and created date/time.`,
	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("list called")
      if len(args) == 0 {
				 fmt.Println("Enter region,refer \n https://docs.aws.amazon.com/general/latest/gr/rande.html")
	       os.Exit(0)
      }

		sess, err := session.NewSession(&aws.Config{
        Region: aws.String(args[0])},
    )

    // Create DynamoDB client
    svc := dynamodb.New(sess)

    // Get the list of tables
    result, err := svc.ListTables(&dynamodb.ListTablesInput{})

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println("Tables:")
//    fmt.Println("")

    for _, n := range result.TableNames {
        fmt.Println(*n)
    }

    fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
