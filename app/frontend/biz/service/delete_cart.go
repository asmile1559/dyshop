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

func (s *DeleteCartService) Run(_ *cart_page.DeleteCartReq) (map[string]interface{}, error) {
	id, ok := s.ctx.Value("user_id").(uint32)
	if !ok {
		return nil, errors.New("no user id")
	}

	resp, err := rpcclient.CartClient.DeleteCart(s.ctx, &pbcart.DeleteCartReq{UserId: id})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil
}
