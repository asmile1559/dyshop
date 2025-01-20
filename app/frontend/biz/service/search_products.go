package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
)

//	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
//	pbbackend "github.com/asmile1559/dyshop/pb/backend/product"

type SearchProductService struct {
	ctx context.Context
}

func NewSearchProductService(c context.Context) *SearchProductService {
	return &SearchProductService{ctx: c}
}

func (s *SearchProductService) Run(req *product_page.SearchProductsReq) (map[string]interface{}, error) {
	//resp, err := rpcclient.ProductClient.SearchProducts(s.ctx, &pbbackend.SearchProductsReq{Query: req.GetQ()})
	//if err != nil {
	//	return nil, err
	//}
	//
	//return gin.H{
	//	"items": resp.Results,
	//	"q": req.GetQ(),
	//}, nil

	return gin.H{
		"status": "search_product ok",
	}, nil
}
