package models

import (
	"log"

	"github.com/DivyanshuVerma98/url-shortener/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var svc *dynamodb.DynamoDB

func init() {
	// Create a new session using the default AWS configuration
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create a DynamoDB client
	svc = dynamodb.New(sess, aws.NewConfig().WithRegion("ap-south-1"))
}

func CreateUrlMapperItem(item *UrlMapperItem) bool {
	log.Println("Inside CreateUrlMapperItem")
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling item: %s", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(config.MappingTableName),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
	return true
}

func GetUrlMapperItem(code string) UrlMapperItem {
	log.Println("Inside GetUrlMapperItem")
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(config.MappingTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"short_url": {
				S: aws.String(code),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem %s", err)
	}
	item := UrlMapperItem{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		log.Fatalf("Got error unmarshalling item: %s", err)
	}
	return item
}
