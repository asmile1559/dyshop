package service

import (
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"golang.org/x/net/context"
)

type ListOrdersService struct {
	ctx context.Context
}

func NewListOrdersService(c context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: c}
}

func (s *ListOrdersService) Run(req *pborder.ListOrderReq) (*pborder.ListOrderResp, error) {
	// TODO: finish your business code...
	//
	return &pborder.ListOrderResp{
		Orders: []*pborder.Order{
			{
				OrderItems: []*pborder.OrderItem{
					{
						Item: &pbcart.CartItem{
							ProductId: 1,
							Quantity:  10,
						},
						Cost: 1000.5,
					},
				},
				OrderId:      "1",
				UserId:       1,
				UserCurrency: "CNY",
				Address: &pborder.Address{
					StreetAddress: "BigStreet",
					City:          "Shenyang",
					State:         "Liaoning",
					Country:       "China",
					ZipCode:       "123456",
				},
				Email:     "123@abc.com",
				CreatedAt: 1312312,
			},
		},
	}, nil

}
