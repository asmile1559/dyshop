package service

import (
	"context"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/gin-gonic/gin"
)

type ListProductService struct {
	Ctx context.Context
}

func NewListProductService(c context.Context) *ListProductService {
	return &ListProductService{
		Ctx: c,
	}
}

func (s *ListProductService) Run(req *pbproduct.ListProductsReq) (map[string]interface{}, error) {
	productClient, conn, err := rpcclient.GetProductClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := productClient.ListProducts(s.Ctx, &pbproduct.ListProductsReq{
		Page:         req.Page,
		PageSize:     req.GetPageSize(),
		CategoryName: req.GetCategoryName(),
	})

	if err != nil {
		return nil, err
	}
	return gin.H{
		"resp": resp,
	}, nil

}
