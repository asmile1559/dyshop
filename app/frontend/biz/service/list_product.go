package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
)

// rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
// pbproduct "github.com/asmile1559/dyshop/pb/backend/product"

type ListProductService struct {
	ctx context.Context
}

func NewListProductService(c context.Context) *ListProductService {
	return &ListProductService{
		ctx: c,
	}
}

func (s *ListProductService) Run(req *product_page.ListProductsReq) (map[string]interface{}, error) {

	//resp, err := rpcclient.ProductClient.ListProducts(s.ctx, &pbproduct.ListProductsReq{
	//	Page:         req.Page,
	//	PageSize:     req.GetPageSize(),
	//	CategoryName: req.GetCategoryName(),
	//})
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return gin.H{
	//	"resp": resp,
	//}, nil

	return gin.H{
		"status": "list cart ok",
	}, nil

}
