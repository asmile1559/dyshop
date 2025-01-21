package main

import (
	service "github.com/asmile1559/dyshop/app/product/biz/service"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"golang.org/x/net/context"
)

type ProductServiceServer struct {
	pbproduct.UnimplementedProductCatalogServiceServer
}

func (s *ProductServiceServer) ListProducts(ctx context.Context, req *pbproduct.ListProductsReq) (*pbproduct.ListProductsResp, error) {
	resp, err := service.NewListProductService(ctx).Run(req)

	return resp, err
}
func (s *ProductServiceServer) GetProduct(ctx context.Context, req *pbproduct.GetProductReq) (*pbproduct.GetProductResp, error) {
	resp, err := service.NewGetProductService(ctx).Run(req)

	return resp, err
}
func (s *ProductServiceServer) SearchProducts(ctx context.Context, req *pbproduct.SearchProductsReq) (*pbproduct.SearchProductsResp, error) {
	resp, err := service.NewSearchProductService(ctx).Run(req)

	return resp, err
}
