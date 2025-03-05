package service

import (
	"context"
	"fmt"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
)

type GetProductService struct {
	ctx context.Context
}

func NewGetProductService(c context.Context) *GetProductService {
	return &GetProductService{ctx: c}
}

func (s *GetProductService) Run(req *product_page.GetProductReq) (gin.H, error) {
	id, ok := s.ctx.Value("user_id").(int64)
	if !ok {
		return nil, fmt.Errorf("no user id in context")
	}

	userClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rpcReq2 := &pbuser.GetUserInfoReq{
		UserId: int64(id),
	}
	respUser, err := userClient.GetUserInfo(s.ctx, rpcReq2)
	if err != nil {
		return nil, err
	}

	productClient, conn, err := rpcclient.GetProductClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	respProduct, err := productClient.GetProduct(s.ctx, &pbproduct.GetProductReq{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	// 转换产品数据

	return gin.H{
		"PageRouter": PageRouter,
		"UserInfo": gin.H{
			"UserId": "",
			"Name":   respUser.GetName(),
		},
		"Product": gin.H{
			"ProductId":   respProduct.Product.Id,
			"ProductImg":  respProduct.Product.Picture,
			"ProductName": respProduct.Product.Name,
			"ProductDesc": respProduct.Product.Description,
			"ProductSpecs": []gin.H{
				{"SpecName": respProduct.Product.Name, "SpecPrice": respProduct.Product.Price, "SpecStock": "120"},
				{"SpecName": respProduct.Product.Name, "SpecPrice": respProduct.Product.Price, "SpecStock": "130"},
				{"SpecName": respProduct.Product.Name, "SpecPrice": respProduct.Product.Price, "SpecStock": "140"},
			},
			"ProductCategories": respProduct.Product.Categories,
			"ProductParams":     []gin.H{
				// {"ParamName": "味道", "ParamValue": "盲盒味"},
				// {"ParamName": "材料", "ParamValue": "面粉，小麦"},
				// {"ParamName": "颜色", "ParamValue": "棕黄色"},
				// {"ParamName": "适宜人群", "ParamValue": "所有年龄人群"},
				// {"ParamName": "定位", "ParamValue": "高端产品"},
			},
			"ProductInsurance": "退货险",
			"ProductExpress":   "包邮",
		},
	}, nil
}
