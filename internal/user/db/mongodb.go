package db

import (
	"context"
	"firstGOProject/internal/user"
	"firstGOProject/pkg/logging"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}

func (d *db) FindUser(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectId. hex: %v", id)
	}
	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return u, fmt.Errorf("failed to find user by id. id: %s. due to error: %v", id, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user by id. id: %s. due to error: %v", id, err)
	}

	return u, nil
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user. error: %s", err)
	}

	d.logger.Debug("convert InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}

	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (d *db) Update(ctx context.Context, user user.User) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
