package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB interface {
	Upsert(ctx context.Context, collectionName string, filter, update any) error
	Drop(ctx context.Context, collectionName string) error
	Close(ctx context.Context) error
}

var _ DB = (*MongoDB)(nil)

type MongoDB struct {
	client *mongo.Client
	dbName string
}

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

func (m *MongoDB) Upsert(ctx context.Context, collectionName string, filter, update any) error {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	opts := options.Update().SetUpsert(true)

	_, err := collection.UpdateOne(ctx, filter, bson.M{"$set": update}, opts)
	if err != nil {
		return fmt.Errorf("failed to upsert record: %w", err)
	}

	return nil
}

func (m *MongoDB) Drop(ctx context.Context, collectionName string) error {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	if err := collection.Drop(ctx); err != nil {
		return fmt.Errorf("failed to drop collection %s.%s: %w", collection.Database().Name(), collection.Name(), err)
	}

	return nil
}

func (m *MongoDB) Close(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
