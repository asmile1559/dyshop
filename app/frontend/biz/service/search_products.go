package service

import (
	"context"
	"fmt"
	"strconv"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
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

	// 调用RPC服务
	productClient, conn, err := rpcclient.GetProductClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	products := []gin.H{}
	if req.GetQuery() == "" {
		rpcReq := &pbproduct.ListProductsReq{
			CategoryName: req.GetCategory(),
			Page:         int32(page),
			PageSize:     int64(pageSize),
		}
		rpcResp, err := productClient.ListProducts(s.ctx, rpcReq)
		if err != nil {
			return nil, err
		}

		for _, p := range rpcResp.Products {
			product := gin.H{
				"ProductId":         p.Id,
				"uid":               p.Uid,
				"ProductImg":        p.Picture,
				"ProductName":       p.Name,
				"ProductDesc":       p.Description,
				"ProductSpecs":      []gin.H{},
				"ProductCategories": []string{},
				"ProductParams":     []gin.H{},
				"ProductInsurance":  "退货险",
				"ProductExpress":    "包邮",
				"ProductSpec": gin.H{
					"Name":  p.Description,
					"Price": p.Price,
				},
				"Quantity": strconv.FormatInt(int64(100), 10),
				"Currency": "CNY",
				"Postage":  fmt.Sprintf("%.2f", 50.0),
				"Sold":     strconv.FormatInt(int64(100), 10),
			}
			products = append(products, product)
		}
	} else {
		rpcReq := &pbproduct.SearchProductsReq{
			Query:    req.GetQuery(),
			Page:     int32(page),
			PageSize: int32(pageSize),
		}
		rpcResp, err := productClient.SearchProducts(s.ctx, rpcReq)
		if err != nil {
			return nil, err
		}

		for _, p := range rpcResp.Results {
			product := gin.H{
				"ProductId":         p.Id,
				"uid":               p.Uid,
				"ProductImg":        p.Picture,
				"ProductName":       p.Name,
				"ProductDesc":       p.Description,
				"ProductSpecs":      []gin.H{},
				"ProductCategories": []string{},
				"ProductParams":     []gin.H{},
				"ProductInsurance":  "退货险",
				"ProductExpress":    "包邮",
				"ProductSpec": gin.H{
					"Name":  p.Description,
					"Price": p.Price,
				},
				"Quantity": strconv.FormatInt(int64(100), 10),
				"Currency": "CNY",
				"Postage":  fmt.Sprintf("%.2f", 50.0),
				"Sold":     strconv.FormatInt(int64(100), 10),
			}
			products = append(products, product)
		}
	}

	UserClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	id := s.ctx.Value("user_id")
	rpcReq2 := &pbuser.GetUserInfoReq{
		UserId: id.(int64),
	}
	rpcResp2, err := UserClient.GetUserInfo(s.ctx, rpcReq2)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"PageRouter": PageRouter, // 需要实现页面路由获取逻辑
		"UserInfo": gin.H{
			"UserId": rpcResp2.GetUserId(),
			"Name":   rpcResp2.GetName(),
			"Sign":   rpcResp2.GetSign(),
			"Img":    rpcResp2.GetUrl(),
			"Role":   rpcResp2.GetRole(),
		},
		"Products": products,
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
