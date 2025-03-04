package service

import (
	"context"

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

func (s *GetProductService) Run(req *product_page.GetProductReq) (map[string]any, error) {
	productClient, conn, err := rpcclient.GetProductClient()
	id, _ := s.ctx.Value("user_id").(uint32)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	userClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	// 构造RPC请求
	rpcReq2 := &pbuser.GetUserInfoReq{
		UserId: int64(id),
	}
	rpcResp2, err := userClient.GetUserInfo(s.ctx, rpcReq2)
	resp, err := productClient.GetProduct(s.ctx, &pbproduct.GetProductReq{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	// 转换产品数据

	return gin.H{
		"PageRouter": PageRouter,
		"UserInfo": gin.H{
			"Id":   rpcResp2.GetUserId(),
			"Name": rpcResp2.GetName(),
			"Sign": rpcResp2.GetSign(),
			"Img":  rpcResp2.GetUrl(),
			"Role": rpcResp2.GetRole(),
		},
		"Products": gin.H{
			"ProductId":   resp.Product.Id,
			"ProductImg":  resp.Product.Picture,
			"ProductName": resp.Product.Name,
			"ProductDesc": resp.Product.Description,
			"ProductSpecs": []gin.H{
				{"SpecName": resp.Product.Name, "SpecPrice": resp.Product.Price, "SpecStock": "120"},
				{"SpecName": resp.Product.Name, "SpecPrice": resp.Product.Price, "SpecStock": "130"},
				{"SpecName": resp.Product.Name, "SpecPrice": resp.Product.Price, "SpecStock": "140"},
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
		},
	}, nil
}
