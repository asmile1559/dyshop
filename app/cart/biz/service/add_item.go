package service

import (
	"context"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	"github.com/asmile1559/dyshop/app/cart/biz/model"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

type AddItemService struct {
	ctx context.Context
}

func NewAddItemService(c context.Context) *AddItemService {
	return &AddItemService{ctx: c}
}

func (s *AddItemService) Run(req *pbcart.AddItemReq) (*pbcart.AddItemResp, error) {
	err := model.AddOrUpdateCartItem(
		dal.DB,
		uint64(req.UserId),
		uint64(req.Item.ProductId),
		int(req.Item.Quantity),
	)
	if err != nil {
		return nil, err
	}
	return &pbcart.AddItemResp{
		CartId: 0, // 仅占位
	}, nil
}
