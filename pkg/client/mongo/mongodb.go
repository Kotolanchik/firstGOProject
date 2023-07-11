package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {
	var mongoDbUrl string
	var isAuth bool

	if username == "" && password == "" {
		mongoDbUrl = fmt.Sprintf("mongodb://%s:%s", host, port)
		isAuth = false
	} else {
		isAuth = true
		mongoDbUrl = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}

	clientOpt := options.Client().ApplyURI(mongoDbUrl)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clientOpt.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		})
	}

	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB. error: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping to mongoDB. error: %v", err)
	}

	return client.Database(database), nil
}
