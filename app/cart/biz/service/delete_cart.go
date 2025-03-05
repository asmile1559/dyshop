package service

import (
	"context"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	"github.com/asmile1559/dyshop/app/cart/biz/model"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

type DeleteCartService struct {
	ctx context.Context
}

func NewDeleteCartService(c context.Context) *DeleteCartService {
	return &DeleteCartService{ctx: c}
}

func (s *DeleteCartService) Run(req *pbcart.DeleteCartReq) (*pbcart.DeleteCartResp, error) {
	userID := uint64(req.UserId)
	var idsToDelete []uint64
	for _, it := range req.Items {
		// it.id 代表 cart_items表里的主键
		idsToDelete = append(idsToDelete, uint64(it.Id))
	}

	// 调用部分删除
	if err := model.DeleteItems(dal.DB, userID, idsToDelete); err != nil {
		return nil, err
	}

	// 如果你想做 "全部删除" 时，用 model.ClearAllItems(...)

	return &pbcart.DeleteCartResp{}, nil
}
