# dyno CLI Command

This project is to create a dyno cli for interacting with AWS dynamodb. Command line will help in
  1) List table
  2) Create table
  3) Update table
  4) Read table
  5) Delete table
  6) Purge table

# Amazon DynamoDB

Amazon DynamoDB is a fully managed proprietary NoSQL database service that supports key-value and document data structures and is offered by Amazon.com as part of the Amazon Web Services portfolio.

# AWS SDK for Go

aws-sdk-go is the official AWS SDK for the Go programming language.

# Environment Variables for SDK

By default, the SDK detects AWS credentials set in your environment and uses them to sign requests to AWS. That way you don't need to manage credentials in your applications. The SDK looks for credentials in the following environment variables:

AWS_ACCESS_KEY_ID / AWS_SECRET_ACCESS_KEY /AWS_SESSION_TOKEN (optional)

The following examples show how you configure the environment variables.

Linux, OS X, or Unix

$ export AWS_ACCESS_KEY_ID=YOUR_AKID <br />
$ export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY <br />
$ export AWS_SESSION_TOKEN=TOKEN <br />
$ export AWS_REGION=us-east-1 <br />

Windows

C:\> set AWS_ACCESS_KEY_ID=YOUR_AKID <br />
C:\> set AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY <br />
C:\> set AWS_SESSION_TOKEN=TOKEN <br />
C:\> set AWS_REGION=us-east-1 <br />

# Available Commands
  create      Create dynamodb table <br />
  delete      Delete the dynamodb table <br />
  help        Help about any command <br />
  list        List all the tables in the DynamoDB <br />
  purge       Purge the dynamodb table <br />
  read        Read the dynamodb table <br />
  region      Set AWS region <br />
  update      Update the dynamodb table <br />

# GO HELP

If you’re ever stuck without internet access, you can get the documentation running locally
via: godoc -http=:6060
and
pointing your browser to http://localhost:6060
