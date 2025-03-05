package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/asmile1559/dyshop/app/order/biz/model"
	"github.com/asmile1559/dyshop/app/order/utils/db"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GetOrderService struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewGetOrderService(c context.Context) *GetOrderService {
	return &GetOrderService{ctx: c, DB: db.DB}
}

func (s *GetOrderService) Run(req *pborder.GetOrderReq) (*pborder.GetOrderResp, error) {
	order := model.Order{}
	err := s.DB.First(&order, req.OrderId).Error
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	addr := &model.Address{}
	s.DB.Model(&model.Address{}).First(&addr, order.AddressId)
	pids := []uint32{}
	pidsStr := strings.Split(order.ProductIDs, ",")
	for _, pidStr := range pidsStr {
		pid := uint32(0)
		fmt.Sscanf(pidStr, "%d", &pid)
		pids = append(pids, pid)
	}
	resp := &pborder.GetOrderResp{
		Order: &pborder.Order{
			Id:     uint32(order.ID),
			UserId: order.UserId,
			Address: &pborder.Address{
				Id:          uint32(addr.ID),
				UserId:      addr.UserId,
				Street:      addr.Street,
				District:    addr.District,
				City:        addr.City,
				Province:    addr.Province,
				Phone:       addr.Phone,
				Recipient:   addr.Recipient,
				FullAddress: addr.FullAddress,
			},
			Price:      order.Price,
			ProductIds: pids,
		},
	}

	return resp, nil
}
