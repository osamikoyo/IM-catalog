package data

import (
	"context"
	"fmt"
	"time"

	"github.com/osamikoyo/IM-catalog/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct{
	coll *mongo.Collection
	ctx context.Context
}

func New(cfg *config.Config) (*Storage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DatabaseURL))
	if err != nil{
		return nil, fmt.Errorf("cant connect to the mongo: %v", err)
	}

	coll := client.Database("im").Collection("catalog")
	return &Storage{
		coll: coll,
		ctx: ctx,
	}, nil
}