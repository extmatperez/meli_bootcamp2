package product

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/extmatperez/meli_bootcamp2/19_storage3/internal/domain"
)

type RepositoryDynamo interface {
	GetAll(ctx context.Context) ([]*domain.ProductDynamo, error)
	Get(ctx context.Context, id string) (*domain.ProductDynamo, error)
	Store(ctx context.Context, product *domain.ProductDynamo) (*domain.ProductDynamo, error)
	Update(ctx context.Context, product *domain.ProductDynamo) (*domain.ProductDynamo, error)
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

func (receiver *dynamoRepository) GetAll(ctx context.Context) ([]*domain.ProductDynamo, error) {
	result, err := receiver.dynamo.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName: aws.String(receiver.table),
	})

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	var products []*domain.ProductDynamo
	for _, item := range result.Items {
		product, err := domain.DynamoItemToProduct(item)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (receiver *dynamoRepository) Get(ctx context.Context, id string) (*domain.ProductDynamo, error) {
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
	return domain.DynamoItemToProduct(result.Item)
}

func (receiver *dynamoRepository) Store(ctx context.Context, product *domain.ProductDynamo) (*domain.ProductDynamo, error) {
	av, err := dynamodbattribute.MarshalMap(product)

	if err != nil {
		return &domain.ProductDynamo{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(receiver.table),
	}

	_, err = receiver.dynamo.PutItemWithContext(ctx, input)

	if err != nil {
		return &domain.ProductDynamo{}, err
	}

	return product, nil
}

func (receiver *dynamoRepository) Delete(ctx context.Context, id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(receiver.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	}

	_, err := receiver.dynamo.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (receiver *dynamoRepository) Update(ctx context.Context, product *domain.ProductDynamo) (*domain.ProductDynamo, error) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(product.Name),
			},
			":p": {
				N: aws.String(fmt.Sprintf("%f", product.Price)),
			},
			":d": {
				S: aws.String(product.Description),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#name":        aws.String("name"),
			"#price":       aws.String("price"),
			"#description": aws.String("description"),
		},
		TableName: aws.String(receiver.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(product.Id),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #name = :n, #price = :p, #description = :d"),
	}

	_, err := receiver.dynamo.UpdateItem(input)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return &domain.ProductDynamo{}, err
	}

	return product, nil
}
