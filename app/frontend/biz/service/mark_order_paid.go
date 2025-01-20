package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
)

// rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
// pborder "github.com/asmile1559/dyshop/pb/backend/order"

type MarkOrderPaidService struct {
	ctx context.Context
}

func NewMarkOrderPaidService(c context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: c}
}

func (s *MarkOrderPaidService) Run(req *order_page.MarkOrderPaidReq) (map[string]interface{}, error) {
	//id, ok := s.ctx.Value("user_id").(uint32)
	//if ok == false {
	//	return nil, errors.New("expect user id")
	//}
	//
	//_, err := rpcclient.OrderClient.MarkOrderPaid(s.ctx, &pborder.MarkOrderPaidReq{
	//	UserId:  id,
	//	OrderId: req.GetOrderId(),
	//})
	//
	//if err != nil {
	//	return nil, err
	//}

	return gin.H{
		"status": "mark_order_paid ok",
	}, nil
}
