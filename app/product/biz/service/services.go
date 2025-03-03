package service

import (
	"context"
	"errors"
	"github.com/asmile1559/dyshop/app/product/biz/dal"
	"github.com/asmile1559/dyshop/app/product/biz/model"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// CreateProductService 处理产品创建
type CreateProductService struct {
	ctx context.Context
}

//
//func init() {
//	dsn := viper.GetString("mysql.dsn")
//	print(dsn)
//	if err := dal.InitDB(dsn); err != nil {
//		logrus.Fatalf("failed to init db: %v", err)
//	}
//}

func NewCreateProductService(c context.Context) *CreateProductService {
	return &CreateProductService{ctx: c}
}

func (s *CreateProductService) Run(req *pbproduct.CreateProductReq) (*pbproduct.CreateProductResp, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "product name is required")
	}
	if req.Price <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid product price")
	}

	newProduct := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  req.Categories,
	}

	err := model.CreateOrUpdateProduct(dal.DB, newProduct)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create product")
	}

	return &pbproduct.CreateProductResp{
		Success: true,
	}, nil
}

// ModifyProductService 处理产品更新
type ModifyProductService struct {
	ctx context.Context
}

func NewModifyProductService(c context.Context) *ModifyProductService {
	return &ModifyProductService{ctx: c}
}

func (s *ModifyProductService) Run(req *pbproduct.ModifyProductReq) (*pbproduct.ModifyProductResp, error) {
	//
	////if req.Id == 0 {
	//	return nil, status.Error(codes.InvalidArgument, "product ID is required")
	//}
	//
	//// 验证至少有一个字段需要更新
	//if req.Name == nil && req.Description == nil && req.Picture == nil && req.Price == nil && req.Categories == nil {
	//	return nil, status.Error(codes.InvalidArgument, "no fields provided for update")
	//}
	//
	newProduct := &model.Product{
		ID:          uint(req.GetId()),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Picture:     req.GetPicture(),
		Price:       req.GetPrice(),
		Categories:  req.Categories,
	}

	err := model.CreateOrUpdateProduct(dal.DB, newProduct)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to update product")
	}

	return &pbproduct.ModifyProductResp{Success: true}, nil
}

// DeleteProductService 处理产品删除
type DeleteProductService struct {
	ctx context.Context
}

func NewDeleteProductService(c context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: c}
}

func (s *DeleteProductService) Run(req *pbproduct.DeleteProductReq) (*pbproduct.DeleteProductResp, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "product ID is required")
	}

	if err := model.DeleteProduct(dal.DB, req.Id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "product not found")
		}
		return nil, status.Error(codes.Internal, "failed to delete product")
	}

	return &pbproduct.DeleteProductResp{Success: true}, nil
}
