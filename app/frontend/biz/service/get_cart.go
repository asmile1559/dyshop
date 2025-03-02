package service

import (
	"context"
	"errors"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"github.com/asmile1559/dyshop/pb/frontend/cart_page"
	"github.com/gin-gonic/gin"
)

type GetCartService struct {
	ctx context.Context
}

func NewGetCartService(c context.Context) *GetCartService {
	return &GetCartService{ctx: c}
}

func (s *GetCartService) Run(_ *cart_page.GetCartReq) (map[string]interface{}, error) {
	id, ok := s.ctx.Value("user_id").(uint32)
	if !ok {
		return nil, errors.New("no user id in context")
	}

	cartClient, conn, err := rpcclient.GetCartClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	resp, err := cartClient.GetCart(s.ctx, &pbcart.GetCartReq{
		UserId: id,
	})
	if err != nil {
		return nil, err
	}

	return gin.H{
		"resp": resp,
	}, nil
}
