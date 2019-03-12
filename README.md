# dyno CLI Command

This project is to create a dyno cli for interacting with AWS dynamodb. Command line will help in
  1) List table
  2) Create table
  3) Update table
  4) Read table
  5) Delete table

# Amazon DynamoDB

Amazon DynamoDB is a fully managed proprietary NoSQL database service that supports key-value and document data structures and is offered by Amazon.com as part of the Amazon Web Services portfolio.

# AWS SDK for Go v2

aws-sdk-go-v2 is the Developer Preview for the v2 of the AWS SDK for the Go programming language.

# Environment Variables for SDK

By default, the SDK detects AWS credentials set in your environment and uses them to sign requests to AWS. That way you don't need to manage credentials in your applications. The SDK looks for credentials in the following environment variables:

AWS_ACCESS_KEY_ID / AWS_SECRET_ACCESS_KEY /AWS_SESSION_TOKEN (optional)

The following examples show how you configure the environment variables.

Linux, OS X, or Unix

$ export AWS_ACCESS_KEY_ID=YOUR_AKID

$ export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY

$ export AWS_SESSION_TOKEN=TOKEN

Windows

$ export AWS_ACCESS_KEY_ID=YOUR_AKID

$ export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY

$ export AWS_SESSION_TOKEN=TOKEN
