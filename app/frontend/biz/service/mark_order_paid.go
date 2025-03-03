package service

import (
	"context"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
)

type MarkOrderPaidService struct {
	ctx context.Context
}

func NewMarkOrderPaidService(c context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: c}
}

func (s *MarkOrderPaidService) Run(req *order_page.MarkOrderPaidReq) (map[string]interface{}, error) {
	orderClient, conn, err := rpcclient.GetOrderClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	_, err = orderClient.MarkOrderPaid(s.ctx, &pborder.MarkOrderPaidReq{
		OrderId: req.GetOrderId(),
	})

	if err != nil {
		return nil, err
	}

	return gin.H{
		"status": "mark_order_paid ok",
	}, nil
}
