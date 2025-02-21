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

func (s *Storage) GetMore(name string) ([]models.Product, error){
	var products []models.Product

	filter, err := s.coll.Find(s.ctx, bson.M{
		"name" : name,
	})
	if err != nil{
		return nil, err
	}

	var product models.Product

	for filter.Next(s.ctx) {
		filter.Decode(&product)

		products = append(products, product)
	}

	return products, nil
}

func (s *Storage) GetAll(name string) ([]models.Product, error) {
	var products []models.Product

	cursor, err := s.coll.Find(s.ctx, bson.M{})
	if err != nil{
		return nil, err
	}

	var product models.Product

	for cursor.Next(s.ctx) {
		if err = cursor.Decode(&product);err != nil{
			fmt.Println(err)
		}

		products = append(products, product)
	}

	if cursor.Err() != nil{
		return nil, err
	}

	return products, nil
}

func (s *Storage) Delete(id uint64) error {
	filter := bson.M{
		"id" : id,
	}

	res := s.coll.FindOneAndDelete(s.ctx, filter)
	return res.Err()
}

func (s *Storage) Update(id uint64, newparams *models.Product) error {
	filter := bson.M{
		"id" : id,
	}

	update, err := bson.Marshal(newparams)
	if err != nil{
		return err
	}

	_, err = s.coll.UpdateOne(s.ctx, filter, update)
	return err
}