package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/gin-gonic/gin"
)

type GetProductService struct {
	ctx context.Context
}

func NewGetProductService(c context.Context) *GetProductService {
	return &GetProductService{ctx: c}
}

func (s *GetProductService) Run(req *product_page.GetProductReq) (map[string]interface{}, error) {
	resp, err := rpcclient.ProductClient.GetProduct(s.ctx, &pbproduct.GetProductReq{Id: req.GetId()})
	if err != nil {
		return nil, err
	}
	return gin.H{

		"ProductId":   resp.Product.Id,
		"ProductImg":  resp.Product.Picture,
		"ProductName": resp.Product.Name,
		"ProductDesc": resp.Product.Description,
		"ProductSpecs": []gin.H{
			{"SpecName": "500g装", "SpecPrice": "18.80", "SpecStock": "120"},
			{"SpecName": "1000g装", "SpecPrice": "36.80", "SpecStock": "130"},
			{"SpecName": "2000g装", "SpecPrice": "72.80", "SpecStock": "140"},
		},
		"ProductCategories": resp.Product.Categories,
		"ProductParams": []gin.H{
			{"ParamName": "味道", "ParamValue": "盲盒味"},
			{"ParamName": "材料", "ParamValue": "面粉，小麦"},
			{"ParamName": "颜色", "ParamValue": "棕黄色"},
			{"ParamName": "适宜人群", "ParamValue": "所有年龄人群"},
			{"ParamName": "定位", "ParamValue": "高端产品"},
		},
		"ProductInsurance": "退货险",
		"ProductExpress":   "包邮",
	}, nil
}
