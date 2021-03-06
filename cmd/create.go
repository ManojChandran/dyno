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
//	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"encoding/json"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create dynamodb table",
	Long: `Create basic dynamodb table with primary and sort key
For example:
dyno create <table name>
`,
	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("create called")
			if len(args) == 0 {
				 fmt.Println("Table Name missing")
				 os.Exit(0)
			}
			query,_ := loadQuery("Query.json")
//			fmt.Println(query)
//  		av, err := dynamodbattribute.ConvertToList(query)
			sess, err := session.NewSession(&aws.Config{
				Region: aws.String(os.Getenv("AWS_REGION"))},
			)

			// Create DynamoDB client
			svc := dynamodb.New(sess)
			input := &dynamodb.CreateTableInput{
			    AttributeDefinitions: []*dynamodb.AttributeDefinition{
			        {
			            AttributeName: aws.String(query.AttributeDefinition[0].AttributeName),
			            AttributeType: aws.String(query.AttributeDefinition[0].AttributeType),
			        },
			        {
			            AttributeName: aws.String(query.AttributeDefinition[1].AttributeName),
			            AttributeType: aws.String(query.AttributeDefinition[1].AttributeType),
			        },
			    },
			    KeySchema: []*dynamodb.KeySchemaElement{
			        {
			            AttributeName: aws.String(query.KeySchema[0].AttributeName),
			            KeyType:       aws.String(query.KeySchema[0].KeyType),
			        },
			        {
			            AttributeName: aws.String(query.KeySchema[1].AttributeName),
			            KeyType:       aws.String(query.KeySchema[1].KeyType),
			        },
			    },
			    ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			        ReadCapacityUnits:  aws.Int64(query.ProvisionedThroughput.ReadCapacityUnits),
			        WriteCapacityUnits: aws.Int64(query.ProvisionedThroughput.WriteCapacityUnits),
			    },
			    TableName: aws.String(args[0]),
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

type Query struct {
  TableName string `json:TableName`
  AttributeDefinition [] struct {
    AttributeName string `json:AttributeDefinition`
    AttributeType string `json:AttributeType`
  }`json:AttributeDefinition`
  KeySchema [] struct {
    AttributeName string `json:AttributeName`
    KeyType string `json:KeyType`
  }`json:KeySchema`
  ProvisionedThroughput struct {
    ReadCapacityUnits int64 `json:ReadCapacityUnits`
    WriteCapacityUnits int64 `json:WriteCapacityUnits`
  }`json:ProvisionedThroughput`
} // Query

func loadQuery(filename string) (Query, error)  {
  var query Query
  queryFile,err := os.Open(filename)
  defer queryFile.Close()
  if err !=nil {
		fmt.Println("--> Query.json missing")
		os.Exit(1)
//    return query, err
  }
  jsonParser := json.NewDecoder(queryFile)
  err = jsonParser.Decode(&query)
  return query, nil
} // loadQuery
