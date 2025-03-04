package main

import (
	"github.com/asmile1559/dyshop/app/product/biz/dal"
	service "github.com/asmile1559/dyshop/app/product/biz/service"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/gin-gonic/gin"
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

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *pbproduct.CreateProductReq) (*pbproduct.CreateProductResp, error) {
	resp, err := service.NewCreateProductService(ctx).Run(req)
	return resp, err
}

func (s *ProductServiceServer) ModifyProduct(ctx context.Context, req *pbproduct.ModifyProductReq) (*pbproduct.ModifyProductResp, error) {
	resp, err := service.NewModifyProductService(ctx).Run(req)
	return resp, err
}

func (s *ProductServiceServer) DeleteProduct(ctx context.Context, req *pbproduct.DeleteProductReq) (*pbproduct.DeleteProductResp, error) {
	resp, err := service.NewDeleteProductService(ctx).Run(req)
	return resp, err
}

// handler.go
func HealthCheck(c *gin.Context) {
	if dal.DB == nil {
		c.JSON(500, gin.H{"status": "database not initialized"})
		return
	}

	sqlDB, _ := dal.DB.DB()
	if err := sqlDB.Ping(); err != nil {
		c.JSON(500, gin.H{"status": "database ping failed"})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
