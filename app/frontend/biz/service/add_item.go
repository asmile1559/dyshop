package service

import (
	"context"
	"errors"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"github.com/asmile1559/dyshop/pb/frontend/cart_page"
	"github.com/gin-gonic/gin"
)

type AddItemService struct {
	ctx context.Context
}

func NewAddItemService(c context.Context) *AddItemService {
	return &AddItemService{ctx: c}
}

func (s *AddItemService) Run(req *cart_page.AddItemReq) (map[string]interface{}, error) {
	id, ok := s.ctx.Value("user_id").(uint32)
	if !ok {
		return nil, errors.New("no user id")
	}

	resp, err := rpcclient.CartClient.AddItem(s.ctx, &pbcart.AddItemReq{
		UserId: id,
		Item: &pbcart.CartItem{
			ProductId: req.GetProductId(),
			Quantity:  req.GetQuantity(),
		},
	})

	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil

	//return gin.H{
	//	"status": "add_cart ok",
	//}, nil
}
