package server

import (
	"context"
	"net/http"

	"github.com/osamikoyo/IM-catalog/internal/data"
	"github.com/osamikoyo/IM-catalog/internal/data/models"
	"github.com/osamikoyo/IM-catalog/pkg/proto/pb"
)

type Server struct{
	pb.UnimplementedCatalogServiceServer
	Storage *data.Storage
}

func (s *Server) Add(_ context.Context, req *pb.AddReq) (*pb.Response, error){
	err := s.Storage.Add(models.ToModels(req.Product))
	if err != nil{
		return &pb.Response{
			Status: http.StatusInternalServerError,
			Error: err.Error(),
		}, nil
	}

	return &pb.Response{
		Status: http.StatusCreated,
		Error: "",
	}, nil
}

func (s *Server) Delete(_ context.Context, req *pb.DeleteReq) (*pb.Response, error){
	err := s.Storage.Delete(req.ID)
	if err != nil{
		return &pb.Response{
			Status: http.StatusInternalServerError,
			Error: err.Error(),
		}, nil
	}

	return &pb.Response{
		Status: http.StatusOK,
		Error: "",
	}, nil
}

func (s *Server) GetAll(_ context.Context, req *pb.GetAllReq) (*pb.GetAllResp, error){
	products, err := s.Storage.GetAll(req.Name)
	if err != nil{
		return &pb.GetAllResp{
			Response: &pb.Response{
				Error: err.Error(),
				Status: http.StatusInternalServerError,
			},
			Products: nil,
		}, nil
	}

	return &pb.GetAllResp{
		Products: models.ToPbArray(products),
		Response: &pb.Response{
			Error: "",
			Status: http.StatusOK,
		},
	}, nil	
}

func (s *Server) GetMore(_ context.Context, req *pb.GetMoreReq) (*pb.GetMoreResp, error){
	products, err := s.Storage.GetAll(req.Name)
	if err != nil{
		return &pb.GetMoreResp{
			Response: &pb.Response{
				Error: err.Error(),
				Status: http.StatusInternalServerError,
			},
			Products: nil,
		}, nil
	}

	return &pb.GetMoreResp{
		Products: models.ToPbArray(products),
		Response: &pb.Response{
			Error: "",
			Status: http.StatusOK,
		},
	}, nil
}
func (s *Server) GetOne(_ context.Context, req *pb.GetOneReq) (*pb.GetOneResp, error){
	product, err := s.Storage.Get(req.ID)
	if err != nil {
		return &pb.GetOneResp{
			Product: nil,
			Response: &pb.Response{
				Error: err.Error(),
				Status: http.StatusInternalServerError,
			},
		}, err
	}

	return &pb.GetOneResp{
		Response: &pb.Response{
			Status: http.StatusOK,
			Error: "",
		},
		Product: models.ToPB(product),
	}, nil
}
func (s *Server) Update(_ context.Context, req *pb.UpdateReq) (*pb.Response, error){
	err := s.Storage.Update(req.ID, models.ToModels(req.NewParams))
	if err != nil{
		return &pb.Response{
			Status: http.StatusInternalServerError,
			Error: err.Error(),
		}, nil
	}

	return &pb.Response{
		Status: http.StatusOK,
		Error: "",
	}, nil
}
