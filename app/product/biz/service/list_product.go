package service

import (
	"context"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
)

type ListProductService struct {
	ctx context.Context
}

func NewListProductService(c context.Context) *ListProductService {
	return &ListProductService{
		ctx: c,
	}
}

func (s *ListProductService) Run(req *pbproduct.ListProductsReq) (*pbproduct.ListProductsResp, error) {
	// TODO: finish your business code...
	//
	return &pbproduct.ListProductsResp{Products: []*pbproduct.Product{
		{
			Id:          1,
			Name:        "haoguozhi",
			Description: "a type of drink",
			Picture:     "https://ss0.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1899428878,3492225815&fm=253&gp=0.jpg",
			Price:       100.0,
			Categories: []string{
				"drink", "daoge",
			},
		},
	}}, nil
}
