package service

import (
	"context"
	"errors"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"github.com/asmile1559/dyshop/pb/frontend/cart_page"
	"github.com/gin-gonic/gin"
)

type EmptyCartService struct {
	ctx context.Context
}

func NewEmptyCartService(c context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: c}
}

func (s *EmptyCartService) Run(_ *cart_page.EmptyCartReq) (map[string]interface{}, error) {

	id, ok := s.ctx.Value("user_id").(uint32)
	if !ok {
		return nil, errors.New("no user id")
	}

	cartClient, conn, err := rpcclient.GetCartClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := cartClient.EmptyCart(s.ctx, &pbcart.EmptyCartReq{UserId: id})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil
}
