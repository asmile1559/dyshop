package service

import (
	"context"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

type GetCartService struct {
	ctx context.Context
}

func NewGetCartService(c context.Context) *GetCartService {
	return &GetCartService{ctx: c}
}

func (s *GetCartService) Run(req *pbcart.GetCartReq) (*pbcart.GetCartResp, error) {
	cart, err := dal.GetCartByUserID(uint64(req.UserId))
	if err != nil {
		return nil, err
	}
	if cart == nil {
		// 用户还没有 cart，返回空
		return &pbcart.GetCartResp{
			Cart: &pbcart.Cart{
				UserId: req.UserId,
				Items:  []*pbcart.CartItem{},
			},
		}, nil
	}

	// 将 CartItems 转成 protobuf 里的 repeated CartItem
	pbItems := make([]*pbcart.CartItem, 0, len(cart.CartItems))
	for _, it := range cart.CartItems {
		pbItems = append(pbItems, &pbcart.CartItem{
			ProductId: uint32(it.ProductId),
			Quantity:  int32(it.Quantity),
		})
	}

	return &pbcart.GetCartResp{
		Cart: &pbcart.Cart{
			UserId: uint32(cart.UserId),
			Items:  pbItems,
		},
	}, nil
}
