package service

import (
	"context"
	"errors"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
)

type ListOrdersService struct {
	ctx context.Context
}

func NewListOrdersService(c context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: c}
}

func (s *ListOrdersService) Run(_ *order_page.ListOrdersReq) (map[string]interface{}, error) {

	id, ok := s.ctx.Value("user_id").(uint32)
	if ok == false {
		return nil, errors.New("expect user id")
	}

	//resp, err := rpcclient.OrderClient.ListOrder(s.ctx, &pborder.ListOrderReq{UserId: id})
	orderClient, conn, err := rpcclient.GetOrderClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := orderClient.ListOrder(s.ctx, &pborder.ListOrderReq{UserId: id})

	if err != nil {
		return nil, err
	}

	return gin.H{
		"orders": resp.GetOrders(),
	}, nil

	//return gin.H{
	//	"status": "list orders ok",
	//}, nil
}
