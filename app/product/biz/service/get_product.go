package service

import (
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"golang.org/x/net/context"
)

type GetProductService struct {
	ctx context.Context
}

func NewGetProductService(c context.Context) *GetProductService {
	return &GetProductService{ctx: c}
}

func (s *GetProductService) Run(req *pbproduct.GetProductReq) (*pbproduct.GetProductResp, error) {
	// TODO: finish your business code...
	//
	return &pbproduct.GetProductResp{
		Product: &pbproduct.Product{
			Id:          1,
			Name:        "haoguozhi",
			Description: "a type of drink",
			Picture:     "https://ss0.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1899428878,3492225815&fm=253&gp=0.jpg",
			Price:       100.0,
			Categories: []string{
				"drink", "daoge",
			},
		},
	}, nil
}
