package service

import (
	"context"
	"errors"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/pb/frontend/cart_page"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GetCartService struct {
	ctx context.Context
}

func NewGetCartService(c context.Context) *GetCartService {
	return &GetCartService{ctx: c}
}

func (s *GetCartService) Run(_ *cart_page.GetCartReq) ([]gin.H, gin.H, error) {
	id, ok := s.ctx.Value("user_id").(int64)
	if !ok {
		return nil, nil, errors.New("no user id in context")
	}

	cartClient, conn, err := rpcclient.GetCartClient()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	resp, err := cartClient.GetCart(s.ctx, &pbcart.GetCartReq{
		UserId: uint32(id),
	})
	if err != nil {
		return nil, nil, err
	}
	logrus.Debug(resp)

	productClient, conn, err := rpcclient.GetProductClient()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	respList := []gin.H{}
	for _, item := range resp.Items {
		productInfo, err := productClient.GetProduct(s.ctx, &pbproduct.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			return nil, nil, err
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
			"Postage":  "10",
		}
		respList = append(respList, itemMap)
	}
	userClient, conn, err := rpcclient.GetUserClient()
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	userInfo, err := userClient.GetUserInfo(s.ctx, &pbuser.GetUserInfoReq{
		UserId: id,
	})
	if err != nil {
		return nil, nil, err
	}
	userResp := gin.H{
		"Name": userInfo.Name,
	}

	return respList, userResp, nil
}
