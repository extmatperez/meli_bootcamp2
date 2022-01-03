package internal

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/extmatperez/meli_bootcamp2/tree/scerca_nahuel/19_storage3/ProyectoEstructura/internal/models"
)

func itemToProducto(input map[string]*dynamodb.AttributeValue) (*models.ProductoDynamo, error) {
	var item models.ProductoDynamo
	err := dynamodbattribute.UnmarshalMap(input, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

type RepositoryDynamo interface {
	Store(ctx context.Context, model *models.ProductoDynamo) error
	GetOne(ctx context.Context, id string) (*models.ProductoDynamo, error)
	Delete(ctx context.Context, id string) error
}

type dynamoRepository struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

func NewRepositoryDynamo(dynamo *dynamodb.DynamoDB, table string) RepositoryDynamo {
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
	result, err := receiver.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
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
	return itemToProducto(result.Item)
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
