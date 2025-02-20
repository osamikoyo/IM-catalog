package data

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/osamikoyo/IM-catalog/internal/config"
	"github.com/osamikoyo/IM-catalog/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (s *Storage) Add(product *models.Product) error {
	product.ID = rand.Uint64()

	_, err := s.coll.InsertOne(s.ctx, product)
	return err
}

func (s *Storage) Get(id uint64) (*models.Product, error){
	cursor := s.coll.FindOne(s.ctx, bson.M{
		"id" : id,
	})

	var product models.Product

	err := cursor.Decode(&product)	
	if err != nil{
		return nil, err
	}

	return &product, nil
}
