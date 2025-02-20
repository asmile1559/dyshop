package service

import (
	"github.com/asmile1559/dyshop/app/product/biz/dal"
	"github.com/asmile1559/dyshop/app/product/biz/model"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type GetProductService struct {
	ctx context.Context
}

func NewGetProductService(c context.Context) *GetProductService {
	return &GetProductService{ctx: c}
}

func (s *GetProductService) Run(req *pbproduct.GetProductReq) (*pbproduct.GetProductResp, error) {
	// TODO: finish your business code...
	var dbProduct model.Product

	// 使用 uint 类型查询
	if err, _ := model.GetProductByID(dal.DB, uint(req.Id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}
	return &pbproduct.GetProductResp{
		Product: dbProduct.ToProto(),
	}, nil

}
