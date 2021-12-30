package product

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/19_storage3/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryMongo interface {
	SetDatabaseAndCollection(database string, collection string) *mongo.Collection
	GetAll(ctx context.Context) ([]domain.ProductMongo, error)
	Get(ctx context.Context, id string) (domain.ProductMongo, error)
	Store(ctx context.Context, product domain.ProductMongo) (domain.ProductMongo, error)
	Update(ctx context.Context, product domain.ProductMongo) (domain.ProductMongo, error)
	Delete(ctx context.Context, id string) error
}

type mongoRepository struct {
	mongoDB    *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoRepository(m *mongo.Client) *mongoRepository {
	return &mongoRepository{
		mongoDB: m,
	}
}

func (m *mongoRepository) SetDatabaseAndCollection(database string, collection string) *mongo.Collection {
	m.database = m.mongoDB.Database(database)
	m.collection = m.database.Collection(collection)
	return m.collection
}

func (m *mongoRepository) GetAll(ctx context.Context) ([]domain.ProductMongo, error) {
	cur, err := m.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var products []domain.ProductMongo
	for cur.Next(ctx) {
		var product domain.ProductMongo
		err := cur.Decode(&product)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (m *mongoRepository) Get(ctx context.Context, id string) (domain.ProductMongo, error) {
	var product domain.ProductMongo
	filter := bson.M{"_id": id}

	err := m.collection.FindOne(ctx, filter).Decode(&product)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		return domain.ProductMongo{}, err
	} else if err != nil {
		return domain.ProductMongo{}, err
	}

	return product, nil
}

func (m *mongoRepository) Store(ctx context.Context, product domain.ProductMongo) (domain.ProductMongo, error) {
	res, err := m.collection.InsertOne(ctx, product)

	if err != nil {
		return domain.ProductMongo{}, err
	}
	id := res.InsertedID
	product.Id = id.(primitive.ObjectID)
	return product, nil
}

func (receiver *mongoRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (receiver *mongoRepository) Update(ctx context.Context, product domain.ProductMongo) (domain.ProductMongo, error) {
	return domain.ProductMongo{}, nil
}
