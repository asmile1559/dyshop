package service

import (
	"context"
	"github.com/asmile1559/dyshop/pb/frontend/cart_page"
	"github.com/gin-gonic/gin"
)

//	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
//	pbbackend "github.com/asmile1559/dyshop/pb/backend/cart"

type AddCartService struct {
	ctx context.Context
}

func NewAddCartService(c context.Context) *AddCartService {
	return &AddCartService{ctx: c}
}

func (s *AddCartService) Run(req *cart_page.AddCartReq) (map[string]interface{}, error) {
	//id, ok := s.ctx.Value("user_id").(uint32)
	//if !ok {
	//	return nil, errors.New("no user id")
	//}

	//_, err = rpcclient.CartClient.AddItem(s.ctx, &pbbackend.AddItemReq{
	//	UserId: id,
	//	Item: &pbbackend.CartItem{
	//		ProductId: req.GetProductId(),
	//		Quantity:  req.GetQuantity(),
	//	},
	//})
	//
	//if err != nil {
	//	return nil, err
	//}

	return gin.H{
		"status": "add_cart ok",
	}, nil
}
