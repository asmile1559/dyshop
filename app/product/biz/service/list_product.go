package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	"github.com/asmile1559/dyshop/app/product/biz/dal"
	"github.com/asmile1559/dyshop/app/product/biz/model"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ListProductService struct {
	ctx context.Context
}

func init() {
	rpcclient.InitRPCClient()
}
func NewListProductService(c context.Context) *ListProductService {
	return &ListProductService{
		ctx: c,
	}
}

func (s *ListProductService) Run(req *pbproduct.ListProductsReq) (*pbproduct.ListProductsResp, error) {
	// TODO: finish your business code...

	if rpcclient.ProductClient == nil {
		return nil, status.Error(codes.Internal, "RPC 客户端未初始化")
	}
	if req.Page < 1 || req.PageSize < 1 {
		return nil, status.Error(codes.InvalidArgument, "分页参数无效")
	}

	// 调用 DAO 层
	products, total, err := model.ListProducts(
		dal.DB,
		req.Page,
		int32(req.PageSize),
		req.CategoryName,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "database error: ")
	}
	print(total)
	// 转换数据到 Protobuf 格式
	pbProducts := make([]*pbproduct.Product, 0, len(products))
	for _, p := range products {
		pbProducts = append(pbProducts, p.ToProto())
	}

	return &pbproduct.ListProductsResp{
		Products: pbProducts,
	}, nil

}
