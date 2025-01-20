package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
)

//	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
//	pbbackend "github.com/asmile1559/dyshop/pb/backend/product"

type GetProductService struct {
	ctx context.Context
}

func NewGetProductService(c context.Context) *GetProductService {
	return &GetProductService{ctx: c}
}

func (s *GetProductService) Run(req *product_page.GetProductReq) (map[string]interface{}, error) {
	//resp, err := rpcclient.ProductClient.GetProduct(s.ctx, &pbbackend.GetProductReq{Id: req.GetId()})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return gin.H{
	//	"item": resp.GetProduct(),
	//}, nil

	return gin.H{
		"status": "get_product ok",
	}, nil
}
