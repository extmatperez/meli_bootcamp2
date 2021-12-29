package internal

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/TM/proyecto/internal/models"
)

//DB
func InitDynamo() (*dynamodb.DynamoDB, error) {

	region := "us-west-2"
	endpoint := "http://localhost:8000"

	cred := credentials.NewStaticCredentials("local", "local", "")
	sess, err := session.NewSession(aws.NewConfig().WithEndpoint(endpoint).WithRegion(region).WithCredentials(cred))

	if err != nil {
		return nil, err
	}

	dynamo := dynamodb.New(sess)

	return dynamo, nil
}

func itemToProductoDynamo(input map[string]*dynamodb.AttributeValue) (*models.ProductoDynamo, error) {

	var item models.ProductoDynamo

	err := dynamodbattribute.UnmarshalMap(input, &item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

//Store
type RepositoryDynamo interface {
	Store(ctx context.Context, model *models.ProductoDynamo) error
	GetOne(ctx context.Context, id string) (*models.ProductoDynamo, error)
	Delete(ctx context.Context, id string) error
}

type dynamoRepository struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

func NewDynamoRepository(dynamo *dynamodb.DynamoDB, table string) RepositoryDynamo {
	return &dynamoRepository{
		dynamo: dynamo,
		table:  table,
	}
}

func (receiver *dynamoRepository) Store(ctx context.Context, model *models.ProductoDynamo) error {

	av, err := dynamodbattribute.MarshalMap(model)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(receiver.table),
	}

	_, err = receiver.dynamo.PutItemWithContext(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

func (receiver *dynamoRepository) GetOne(ctx context.Context, id string) (*models.ProductoDynamo, error) {

	result, err := receiver.dynamo.GetItemWithContext(ctx,
		&dynamodb.GetItemInput{
			TableName: aws.String(receiver.table),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(id),
				},
			},
		})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	return itemToProductoDynamo(result.Item)
}

func (receiver *dynamoRepository) Delete(ctx context.Context, id string) error {

	_, err := receiver.dynamo.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(receiver.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
