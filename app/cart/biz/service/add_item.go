package service

import (
	"context"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

type AddItemService struct {
	ctx context.Context
}

func NewAddItemService(c context.Context) *AddItemService {
	return &AddItemService{ctx: c}
}

func (s *AddItemService) Run(req *pbcart.AddItemReq) (*pbcart.AddItemResp, error) {
	// user_id, product_id, quantity 都是 uint64/int
	err := dal.AddOrUpdateCartItem(
		uint64(req.UserId),
		uint64(req.Item.ProductId),
		int(req.Item.Quantity),
	)
	if err != nil {
		return nil, err
	}
	return &pbcart.AddItemResp{}, nil
}
