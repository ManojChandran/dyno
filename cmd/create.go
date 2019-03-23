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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create dynamodb table",
	Long: `Create basic dynamodb table with primary and sort key
For example:
dyno create us-east-1
`,
	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("create called")
			if len(args) == 0 {
				 fmt.Println("Need region info,refer \n https://docs.aws.amazon.com/general/latest/gr/rande.html")
				 os.Exit(0)
			}

			sess, err := session.NewSession(&aws.Config{
				Region: aws.String(args[0])},
			)

			// Create DynamoDB client
			svc := dynamodb.New(sess)
			input := &dynamodb.CreateTableInput{
			    AttributeDefinitions: []*dynamodb.AttributeDefinition{
			        {
			            AttributeName: aws.String("year"),
			            AttributeType: aws.String("N"),
			        },
			        {
			            AttributeName: aws.String("title"),
			            AttributeType: aws.String("S"),
			        },
			    },
			    KeySchema: []*dynamodb.KeySchemaElement{
			        {
			            AttributeName: aws.String("year"),
			            KeyType:       aws.String("HASH"),
			        },
			        {
			            AttributeName: aws.String("title"),
			            KeyType:       aws.String("RANGE"),
			        },
			    },
			    ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			        ReadCapacityUnits:  aws.Int64(10),
			        WriteCapacityUnits: aws.Int64(10),
			    },
			    TableName: aws.String(args[1]),
			}
			_, err = svc.CreateTable(input)
			if err != nil {
			    fmt.Println("Got error calling CreateTable:")
			    fmt.Println(err.Error())
			    os.Exit(1)
			}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
