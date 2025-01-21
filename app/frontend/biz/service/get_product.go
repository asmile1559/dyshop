package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
)

type GetProductService struct {
	ctx context.Context
}

func NewGetProductService(c context.Context) *GetProductService {
	return &GetProductService{ctx: c}
}

func (s *GetProductService) Run(req *product_page.GetProductReq) (map[string]interface{}, error) {
	resp, err := rpcclient.ProductClient.GetProduct(s.ctx, &pbproduct.GetProductReq{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil

	//return gin.H{
	//	"status": "get_product ok",
	//}, nil
}
