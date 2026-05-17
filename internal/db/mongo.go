package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	dbName string
}

type Filter map[string]any

var ErrNotFound = mongo.ErrNoDocuments

func Connect(uri, dbName string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("Error when attempting to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Error when pinging MongoDB server: %w", err)
	}

	return &MongoDB{client: client, dbName: dbName}, nil
}

func FindOne[T any](ctx context.Context, m *MongoDB, collectionName string, filter Filter) (*T, error) {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	var result T

	err := collection.FindOne(ctx, bson.M(filter)).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("findone %s: %w", collectionName, err)
	}

	return &result, nil
}

func GetAll[T any](ctx context.Context, m *MongoDB, collectionName string) ([]T, error) {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to query %s: %w", collectionName, err)
	}
	defer cursor.Close(ctx)

	var result []T
	if err := cursor.All(ctx, &result); err != nil {
		return nil, fmt.Errorf("failed to decode %s: %w", collectionName, err)
	}

	return result, nil
}

func Upsert(ctx context.Context, m *MongoDB, collectionName string, filter Filter, update any) error {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, bson.M(filter), bson.M{"$set": update}, opts)
	if err != nil {
		return fmt.Errorf("failed to upsert record: %w", err)
	}

	return nil
}

func Drop(ctx context.Context, m *MongoDB, collectionName string) error {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	if err := collection.Drop(ctx); err != nil {
		return fmt.Errorf("failed to drop collection %s.%s: %w", collection.Database().Name(), collection.Name(), err)
	}

	return nil
}

func Close(ctx context.Context, m *MongoDB) error {
	return m.client.Disconnect(ctx)
}
