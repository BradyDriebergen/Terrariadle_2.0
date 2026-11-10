package db

import (
	"context"
	"errors"
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

	// Fix added during 11/3/26 incident. See docs for more information
	// Optimized pool initialization for 1 vCPU environments and strict cloud firewalls
	opts := options.Client().
		ApplyURI(uri).
		// Keep background monitoring threads warm and active
		SetMinPoolSize(2).
		// Limit context-switching overhead on a fractional/shared vCPU instance
		SetMaxPoolSize(15).
		// Prune unused connections cleanly inside our 120s keep-alive window
		SetMaxConnIdleTime(2 * time.Minute).
		// Hard operational deadline stop to prevent application-wide stalling
		SetTimeout(5 * time.Second)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping MongoDB: %w", err)
	}

	return &MongoDB{client: client, dbName: dbName}, nil
}

func FindOne[T any](ctx context.Context, m *MongoDB, collectionName string, filter Filter) (*T, error) {
	collection := m.client.Database(m.dbName).Collection(collectionName)

	var result T

	err := collection.FindOne(ctx, bson.M(filter)).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
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

func DeleteAll(ctx context.Context, m *MongoDB, collectionName string) error {
	collection := m.client.Database(m.dbName).Collection(collectionName)
	if _, err := collection.DeleteMany(ctx, bson.D{}); err != nil {
		return fmt.Errorf("failed to delete all documents from %s.%s: %w", collection.Database().Name(), collection.Name(), err)
	}
	return nil
}

func Close(ctx context.Context, m *MongoDB) error {
	return m.client.Disconnect(ctx)
}
