package service

import (
	"github.com/asmile1559/dyshop/app/product/biz/dal"
	"github.com/asmile1559/dyshop/app/product/biz/model"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	var dbProduct *model.Product
	var err error
	// 使用 uint 类型查询
	if dbProduct, err = model.GetProductByID(dal.DB, uint(req.Id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithField("product_id", req.Id).Warn("product not found")
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		logrus.WithError(err).Error("failed to get product by id")
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}
	return &pbproduct.GetProductResp{
		Product: dbProduct.ToProto(),
	}, nil

}
