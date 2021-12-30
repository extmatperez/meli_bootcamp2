package domain

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ProductDynamo struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ProductMongo struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Price       float64            `json:"price"`
	Description string             `json:"description"`
}

func DynamoItemToProduct(input map[string]*dynamodb.AttributeValue) (*ProductDynamo, error) {
	var item ProductDynamo
	err := dynamodbattribute.UnmarshalMap(input, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
