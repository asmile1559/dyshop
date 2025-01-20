package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
)

// rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
// pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
// pborder "github.com/asmile1559/dyshop/pb/backend/order"

type GetOrderService struct {
	ctx context.Context
}

func NewGetOrderService(c context.Context) *GetOrderService {
	return &GetOrderService{ctx: c}
}

func (s *GetOrderService) Run(req *order_page.GetOrderReq) (map[string]interface{}, error) {

	//id, ok := s.ctx.Value("user_id").(uint32)
	//if ok == false {
	//	return nil, errors.New("expect user id")
	//}
	//
	//// GetOrderReq not defined in protobuf,
	//// it's just a model
	//resp, err := rpcclient.OrderClient.ListOrder(s.ctx, &pborder.GetOrderReq{
	//	UserId:  id,
	//	OrderId: req.GetOrderId(),
	//})
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return gin.H{
	//	"order": resp.GetOrder(),
	//}, nil

	return gin.H{
		"status": "get order ok",
	}, nil
}
