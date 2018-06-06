# Lambda Go VPC Experiment
## Intro
This repository exists to explore the affect of delpoying Lambda functions into a VPC.  
The repository is split up into 3 different stacks
 - vpc_cloudformation_template.yml - A stock standard VPC template that exposes one or two more things than normal
 - vpc.yml - Two lambda functions, an RDS instance encapsulated by a VPC & a cloudwatch rule to trigger the lambda functions
 - novpc.yml - Two lambda functions and an RDS instance & a cloudwatch rule to trigger the lambda functions

</a>
Once these are deployed you will be able to explore the cold start times introduced by a VPC.
I would normally set up the RDS instance manually with tables, but for the VPC encapulated vpc stack I added an init function to create the database and table, called `init`.


## Build

```bash
# Build binary
$ > make build

# Test Go Code
$ > make test
```

## Deploy

### Create .env

```bash
AWS_ACCOUNT_ID=1234567890
AWS_BUCKET_NAME=your-bucket-name-for-cloudformation-package-data

AWS_STACK_NAME_LAMBDA=lambda-go-vpc-experiment
AWS_STACK_NAME_LAMBDA_VPC=lambda-go-vpc-experiment-with-vpc
AWS_STACK_NAME_VPC=lambda-go-vpc-experiment-a-vpc
AWS_REGION=us-west-1
PARAMETER_OVERRIDES=dbUserName=username dbUserPassword=password
PARAMETER_OVERRIDES_VPC="EnvironmentName=lambda-vpc-experiment"
```

### Install AWS CLI

```bash
$ > brew install awscli
```

### Command

```bash
# Create S3 Bucket
$ > make configure

# Upload data to S3 Bucket
$ > make package

# Deploy CloudFormation Stack
$ > make deploy
```

## Usage

```bash
$ > make outputs

[
  {
    "OutputKey": "URL",
    "OutputValue": "https://random-id.execute-api.us-west-1.amazonaws.com/Prod",
    "Description": "URL for HTTPS Endpoint"
  }
]

$ > curl https://random-id.execute-api.us-west-1.amazonaws.com/Stage/people

{"data":[{"id":"d1","name":"Anton","age":31},{"id":"c2","name":"Frank","age":28},{"id":"b1","name":"Horst","age":42}]}

$ > curl https://random-id.execute-api.us-west-1.amazonaws.com/Stage/person/b1

{"data":{"id":"b1","name":"Horst","age":42}}
```