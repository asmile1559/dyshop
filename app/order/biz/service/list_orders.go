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

type ListOrdersService struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewListOrdersService(c context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: c, DB: db.DB}
}

func (s *ListOrdersService) Run(req *pborder.ListOrderReq) (*pborder.ListOrderResp, error) {
	orders := []model.Order{}
	err := s.DB.Find(&orders).Error
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	rpcOrders := []*pborder.Order{}
	for _, order := range orders {
		addr := &model.Address{}
		s.DB.Model(&model.Address{}).First(&addr, order.AddressId)
		pids := []uint32{}
		pidsStr := strings.Split(order.ProductIDs, ",")
		for _, pidStr := range pidsStr {
			pid := uint32(0)
			fmt.Sscanf(pidStr, "%d", &pid)
			pids = append(pids, pid)
		}
		rpcOrder := &pborder.Order{
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
		}
		rpcOrders = append(rpcOrders, rpcOrder)
	}

	return &pborder.ListOrderResp{
		Orders: rpcOrders,
	}, nil
}
