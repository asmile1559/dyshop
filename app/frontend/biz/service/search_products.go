package service

import (
	"context"
	"fmt"
	"strconv"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type SearchProductService struct {
	ctx context.Context
}

func NewSearchProductService(c context.Context) *SearchProductService {
	return &SearchProductService{ctx: c}
}

func (s *SearchProductService) Run(req *product_page.SearchProductsReq) (gin.H, error) {
	// 参数验证
	if req.GetQuery() == "" && req.GetCategory() == "" {
		return nil, errors.New("must have a keyword or a category")
	}

	// 处理分页参数
	page := int(req.GetPage())
	if page <= 0 {
		page = 1
	}

	pageSize := int(req.GetPageSize())
	if pageSize <= 0 {
		pageSize = 30
	}

	// 处理排序参数
	sort := req.GetSort()
	if sort == "" {
		sort = "comprehensive"
	}

	// 构造RPC请求
	rpcReq := &pbproduct.SearchProductsReq{
		Query:    req.GetCategory() + req.GetQuery(),
		Page:     int32(page),
		PageSize: int32(pageSize),
	}

	// 调用RPC服务
	productClient, conn, err := rpcclient.GetProductClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcResp, err := productClient.SearchProducts(s.ctx, rpcReq)
	if err != nil {
		return nil, errors.Wrap(err, "RPC call failed")
	}

	// 转换产品数据
	products := make([]gin.H, 0, len(rpcResp.Results))
	for _, p := range rpcResp.Results {
		product := gin.H{
			"ProductId":   p.Id,
			"ProductImg":  p.Picture,
			"ProductName": p.Name,
			"ProductSpec": gin.H{
				"Name":  p.Description,
				"Price": fmt.Sprintf("%.2f", p.Description),
			},
			"Quantity": strconv.FormatInt(int64(100), 10),
			"Currency": "CNY",
			"Postage":  fmt.Sprintf("%.2f", 50),
			"Sold":     strconv.FormatInt(int64(100), 10),
		}
		products = append(products, product)
	}

	// 构造最终响应
	return gin.H{
		"PageRouter": getPageRouter(), // 需要实现页面路由获取逻辑
		"UserInfo": gin.H{
			"Name": getUserName(s.ctx), // 需要实现用户信息获取逻辑
		},
		"Products":  products,
		"Keyword":   req.GetQuery(),
		"Sort":      sort,
		"CurPage":   page,
		"pageSize":  pageSize,
		"TotalPage": 50,
	}, nil
}

// 示例辅助函数（需要根据项目实际实现）
func getPageRouter() string {
	return "/search"
}

func getUserName(ctx context.Context) string {
	// 从上下文中获取用户信息
	return "lixiaoming"
}
