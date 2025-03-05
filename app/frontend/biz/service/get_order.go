package service

import (
	"context"
	"errors"

	"slices"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GetOrderService struct {
	ctx context.Context
}

func NewGetOrderService(c context.Context) *GetOrderService {
	return &GetOrderService{ctx: c}
}

func (s *GetOrderService) Run(req *order_page.GetOrderReq) (gin.H, error) {
	id, ok := s.ctx.Value("user_id").(int64)
	if !ok {
		return nil, errors.New("expect user id")
	}

	// 获取user信息
	userClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	userResp, err := userClient.GetUserInfo(s.ctx, &pbuser.GetUserInfoReq{
		UserId: id,
	})
	if err != nil {
		return nil, err
	}

	orderClient, conn, err := rpcclient.GetOrderClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	orderResp, err := orderClient.GetOrder(s.ctx, &pborder.GetOrderReq{
		OrderId: req.OrderId,
	})
	if err != nil {
		return nil, err
	}

	respAddr := orderResp.Order.GetAddress()
	addr := gin.H{
		"AddressId":   respAddr.GetId(),
		"Recipient":   respAddr.GetRecipient(),
		"Phone":       respAddr.GetPhone(),
		"Province":    respAddr.GetProvince(),
		"City":        respAddr.GetCity(),
		"District":    respAddr.GetDistrict(),
		"Street":      respAddr.GetStreet(),
		"FullAddress": respAddr.GetFullAddress(),
	}

	cartClient, conn, err := rpcclient.GetCartClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cartResp, err := cartClient.GetCart(s.ctx, &pbcart.GetCartReq{
		UserId: uint32(id),
	})
	if err != nil {
		return nil, err
	}

	productClient, conn, err := rpcclient.GetProductClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	pids := orderResp.Order.GetProductIds()
	productList := []gin.H{}
	for _, item := range cartResp.Items {
		logrus.Debug(item.Id, pids)
		if !slices.Contains(pids, uint32(item.Id)) {
			continue
		}

		productInfo, err := productClient.GetProduct(s.ctx, &pbproduct.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			return nil, err
		}
		logrus.Debug(productInfo)
		itemMap := gin.H{
			"ItemId":      item.Id,
			"ProductId":   item.ProductId,
			"ProductImg":  productInfo.Product.Picture,
			"ProductName": productInfo.Product.Name,
			"ProductSpec": gin.H{
				"Name":  "",
				"Price": productInfo.Product.Price,
			},
			"Quantity": item.Quantity,
			"Currency": "CNY",
			"Postage":  "10",
		}
		productList = append(productList, itemMap)
	}

	return gin.H{
		"PageRouter": PageRouter,
		"UserInfo": gin.H{
			"Name": userResp.GetName(),
		},
		"AddressInfo": gin.H{ // 自维护
			"Default":   1,
			"Addresses": []gin.H{addr},
		},
		"Products":        productList,
		"OrderPrice":      orderResp.Order.GetPrice(),
		"OrderPostage":    "10.00",
		"OrderDiscount":   "0",
		"OrderFinalPrice": orderResp.Order.GetPrice() + 10.0,
	}, nil
}
