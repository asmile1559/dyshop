package service

import (
	"context"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	"github.com/asmile1559/dyshop/app/cart/biz/model"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

type GetCartService struct {
	ctx context.Context
}

func NewGetCartService(c context.Context) *GetCartService {
	return &GetCartService{ctx: c}
}

func (s *GetCartService) Run(req *pbcart.GetCartReq) (*pbcart.GetCartResp, error) {
	userID := uint64(req.UserId)

	// 取所有 cart_items
	items, err := model.GetCartItemsByUserID(dal.DB, userID)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		// 没有记录
		return &pbcart.GetCartResp{
			UserId: req.UserId,
			Items:  []*pbcart.CartItem{},
		}, nil
	}

	// 转换为 proto 结构
	var pbItems []*pbcart.CartItem
	for _, it := range items {
		pbItems = append(pbItems, &pbcart.CartItem{
			Id:        int32(it.ID),
			UserId:    int32(it.UID),
			ProductId: uint32(it.ProductId),
			Quantity:  int32(it.Quantity),
		})
	}

	return &pbcart.GetCartResp{
		UserId: req.UserId,
		Items:  pbItems,
	}, nil
}
