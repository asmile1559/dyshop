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
	// 获取用户对应的购物车
	cart := dal.GetCartByUserID(req.UserId)

	// 判断购物车中是否已经存在同一 product_id
	found := false
	for i, item := range cart.Items {
		if item.ProductID == req.Item.ProductId {
			cart.Items[i].Quantity += req.Item.Quantity
			found = true
			break
		}
	}

	// 若未找到相同商品，则追加
	if !found {
		cart.Items = append(cart.Items, model.CartItem{
			ProductID: req.Item.ProductId,
			Quantity:  req.Item.Quantity,
		})
	}

	// 持久化
	dal.SaveCart(cart)

	return &pbcart.AddItemResp{}, nil
}
