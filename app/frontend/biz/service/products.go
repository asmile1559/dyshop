package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
)

// CreateProductService 处理商品创建服务
type CreateProductService struct {
	Ctx context.Context
}

func NewCreateProductService(c context.Context) *CreateProductService {
	return &CreateProductService{Ctx: c}
}

func (s *CreateProductService) Run(req *product_page.CreateProductReq) (map[string]interface{}, error) {
	resp, err := rpcclient.ProductClient.CreateProduct(s.Ctx, &pbproduct.CreateProductReq{
		Name:        req.GetName(),
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	})
	if err != nil {
		return nil, err
	}
	return gin.H{"resp": resp}, nil
}

// ModifyProductService 处理商品修改服务
type ModifyProductService struct {
	c context.Context
}

func NewModifyProductService(c context.Context) *ModifyProductService {
	return &ModifyProductService{c: c}
}

func (s *ModifyProductService) Run(req *product_page.ModifyProductReq) (map[string]interface{}, error) {

	resp, err := rpcclient.ProductClient.ModifyProduct(s.c, &pbproduct.ModifyProductReq{
		Id:          req.GetId(),
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	})
	if err != nil {
		return nil, err
	}
	return gin.H{"resp": resp}, nil
}

// DeleteProductService 处理商品删除服务
type DeleteProductService struct {
	Ctx context.Context
}

func NewDeleteProductService(c context.Context) *DeleteProductService {
	return &DeleteProductService{Ctx: c}
}

func (s *DeleteProductService) Run(req *product_page.DeleteProductReq) (map[string]interface{}, error) {
	resp, err := rpcclient.ProductClient.DeleteProduct(s.Ctx, &pbproduct.DeleteProductReq{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return gin.H{"resp": resp}, nil
}
