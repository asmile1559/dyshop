package service

import (
	"github.com/asmile1559/dyshop/app/product/biz/dal"
	"github.com/asmile1559/dyshop/app/product/biz/model"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
)

type SearchProductService struct {
	ctx context.Context
}

func NewSearchProductService(c context.Context) *SearchProductService {
	return &SearchProductService{ctx: c}
}

func (s *SearchProductService) Run(req *pbproduct.SearchProductsReq) (*pbproduct.SearchProductsResp, error) {
	// 参数处理
	page, pageSize := validatePagination(req.GetPage(), req.GetPageSize())
	Name := req.GetQuery()

	// 执行数据库查询
	products, total, err := model.SearchProducts(dal.DB, page, pageSize, Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "database error: %v", err)
	}

	// 转换Protobuf格式
	pbProducts := make([]*pbproduct.Product, 0, len(products))
	for _, p := range products {
		pbProducts = append(pbProducts, p.ToProto())
	}

	// 计算总页数
	totalPages := int32(math.Ceil(float64(total) / float64(pageSize)))

	return &pbproduct.SearchProductsResp{
		Results:    pbProducts,
		TotalPages: totalPages,
	}, nil //var dbProduct model.Product

	// 使用 uint 类型查询
	//if err, dbProdcut := model.ListProducts(dal.DB, dal.DB, req.Query); err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		return nil, status.Errorf(codes.NotFound, "product not found")
	//	}
	//	return nil, status.Errorf(codes.Internal, "database error: %v", err)
	//}
	pbProducts = make([]*pbproduct.Product, 0, 1)
	return &pbproduct.SearchProductsResp{
		Results: pbProducts, // 转换模型到 Protobuf
	}, nil
	// TODO: finish your business code...
	//

}

// 参数验证和默认值处理
func validatePagination(page, pageSize int32) (int32, int32) {
	const (
		defaultPage     = 1
		defaultPageSize = 30
		maxPageSize     = 100
	)

	if page < 1 {
		page = defaultPage
	}
	if pageSize < 1 || pageSize > maxPageSize {
		pageSize = defaultPageSize
	}
	return page, pageSize
}
