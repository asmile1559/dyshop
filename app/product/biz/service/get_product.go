package service

import (
	"github.com/asmile1559/dyshop/app/product/biz/model"
	"github.com/asmile1559/dyshop/app/product/dao"
	pbproduct "
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"github.com/asmile1559/dyshop/app/product/dao"
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
	if err := dao.Db.First(&dbProduct, uint(req.Id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "product not found")
		}
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}

	return &pbproduct.GetProductResp{
		Product: dbProduct.ToProto(), // 转换模型到 Protobuf
	}, nil

}
