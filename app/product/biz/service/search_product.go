package service

import (
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"golang.org/x/net/context"
)

type SearchProductService struct {
	ctx context.Context
}

func NewSearchProductService(c context.Context) *SearchProductService {
	return &SearchProductService{ctx: c}
}

func (s *SearchProductService) Run(req *pbproduct.SearchProductsReq) (*pbproduct.SearchProductsResp, error) {
	//var dbProduct model.Product

	// 使用 uint 类型查询
	//if err, dbProdcut := model.ListProducts(dal.DB, dal.DB, req.Query); err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return nil, status.Errorf(codes.NotFound, "product not found")
	//	}
	//	return nil, status.Errorf(codes.Internal, "database error: %v", err)
	//}
	pbProducts := make([]*pbproduct.Product, 0, 1)
	return &pbproduct.SearchProductsResp{
		Results: pbProducts, // 转换模型到 Protobuf
	}, nil
	// TODO: finish your business code...
	//

}
