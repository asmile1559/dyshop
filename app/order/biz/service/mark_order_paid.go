package service

import (
	"context"

	"github.com/asmile1559/dyshop/app/order/utils/db"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MarkOrderPaidService struct {
	ctx context.Context
	DB  *gorm.DB
}

func NewMarkOrderPaidService(c context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: c, DB: db.DB}
}

func (s *MarkOrderPaidService) Run(req *pborder.MarkOrderPaidReq) (*pborder.MarkOrderPaidResp, error) {
	logrus.Debug("MarkOrderPaidService.TODO")
	return nil, nil
}
