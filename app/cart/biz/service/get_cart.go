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
	cart := dal.GetCartByUserID(req.UserId)

	// 转换到 proto 定义的 CartItem
	items := make([]*pbcart.CartItem, 0, len(cart.Items))
	for _, i := range cart.Items {
		items = append(items, &pbcart.CartItem{
			ProductId: i.ProductID,
			Quantity:  i.Quantity,
		})
	}

	// 返回给客户端
	return &pbcart.GetCartResp{
		Cart: &pbcart.Cart{
			UserId: cart.UserID,
			Items:  items,
		},
	}, nil
}
