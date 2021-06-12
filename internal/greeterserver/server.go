package greeterserver

import (
	"context"
	"fmt"
	"math/rand"
	"github.com/twitchtv/twirp"
	pb "GreeterService/rpc/greeter"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)


// Server implements the Greeter service
type Server struct{}

//
func (s *Server) SetGreetingForUser(ctx context.Context, name *pb.Name) (empty *pb.Empty, err error) {
	if name.Message == "" {
		return nil, twirp.InvalidArgumentError("Name:", "No name given")
	}
	// dynamo = ConnectDynamo()
	PutItem(name.Message)
	return nil, nil
}

func (s *Server) GetGreetingForUser(ctx context.Context, name *pb.Name) (greeting *pb.Greeting, err error) {
	if name.Message == "" {
		return nil, twirp.InvalidArgumentError("Name:", "No name given")
	}
	var res = GetItem(name.Message)
	return &pb.Greeting{Message: res.message}, nil
}


// DynamoDB connection and methods
type Greeting struct {
	name, message string
}

var dynamo *dynamodb.DynamoDB

func ConnectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("*************", "**********************************", ""),
	})))
}

func init() {
	dynamo = ConnectDynamo()
}

func PutItem(name string) {
	var greet = []string{"Hey there, ", "Hello ", "Hey ", "Nice to meet you, ", "Hi "}[rand.Intn(4)]
	var greeting = greet + name + "!"
	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(name),
			},
			"greeting": {
				S: aws.String(greeting),
			},
		},
		TableName: aws.String("Users"),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func GetItem(name string) (greeting Greeting) {
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(name),
			},
		},
		TableName: aws.String("Users"),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
	err = dynamodbattribute.Unmarshal(result.Item["greeting"], &greeting.message)
	if err != nil {
		panic(err)
	}
	greeting.name = name
	return greeting
}
