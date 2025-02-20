package models

import "github.com/osamikoyo/IM-catalog/pkg/proto/pb"

type Product struct{
	ID uint64 `bson:"id"`
	Name string `bson:"name"`
	Desc string `bson:"desc"`
	Price uint64 `bson:"price"`
}

func ToModels(p *pb.Product) *Product {
	return &Product{
		ID: p.ID,
		Name: p.Name,
		Desc: p.Description,
		Price: p.Price,
	}
}

func ToPB(p *Product) *pb.Product {
	return &pb.Product{
		ID: p.ID,
		Name: p.Name,
		Description: p.Desc,
		Price: p.Price,
	}
}