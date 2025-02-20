package models

type Product struct{
	ID uint64 `bson:"id"`
	Name string `bson:"name"`
	Desc string `bson:"desc"`
	Price uint64 `bson:"price"`
}