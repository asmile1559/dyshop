package service

import (
	"context"
	"errors"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"github.com/asmile1559/dyshop/pb/frontend/cart_page"
	"github.com/gin-gonic/gin"
)

type DeleteCartService struct {
	ctx context.Context
}

func NewDeleteCartService(c context.Context) *DeleteCartService {
	return &DeleteCartService{ctx: c}
}

func (s *DeleteCartService) Run(req *cart_page.DeleteCartReq) (map[string]interface{}, error) {
	id, ok := s.ctx.Value("user_id").(int64)
	if !ok {
		return nil, errors.New("no user id in context")
	}

	cartClient, conn, err := rpcclient.GetCartClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// 把 front-end proto 里的 item_ids 转成后端proto的 repeated CartItem
	var items []*pbcart.CartItem
	for _, itemID := range req.GetItemIds() {
		items = append(items, &pbcart.CartItem{
			Id: int32(itemID),
			// 如果后端需要 user_id/product_id 也可以加
		})
	}

	resp, err := cartClient.DeleteCart(s.ctx, &pbcart.DeleteCartReq{
		UserId: uint32(id),
		Items:  items,
	})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil
}
