package service

import (
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"golang.org/x/net/context"
)

type MarkOrderPaidService struct {
	ctx context.Context
}

func NewMarkOrderPaidService(c context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: c}
}

/*func (s *MarkOrderPaidService) Run(req *pborder.MarkOrderPaidReq) (*pborder.MarkOrderPaidResp, error) {
	// TODO: finish your business code...
	//
	return &pborder.MarkOrderPaidResp{}, nil
}*/

func (s *MarkOrderPaidService) Run(req *pborder.MarkOrderPaidReq) (*pborder.MarkOrderPaidResp, error) {
	// 假设这里有一些业务逻辑处理代码...
	// 例如：更新数据库中的订单状态、发送通知等...

	// 示例数据填充，实际使用时应该替换为从请求或其他服务获取的数据
	updatedOrders := []*pborder.Order{
		{
			OrderId:      req.OrderId,
			UserId:       req.UserId,
			UserCurrency: "USD",      // 用户货币假设为美元
			CreatedAt:    1675680000, // 时间戳示例
			OrderItems: []*pborder.OrderItem{
				{
					Item: &pbcart.CartItem{
						ProductId: 101,
						Quantity:  2,
					},
					Cost: 49.98,
				},
				// 可以添加更多的 OrderItem 实例...
			},
			Address: &pborder.Address{
				StreetAddress: "123 Main St",
				City:          "Anytown",
				State:         "CA",
				Country:       "USA",
				ZipCode:       "90210",
			},
			Email: "user@example.com",
		},
		// 可以添加更多的 Order 实例...
	}

	resp := &pborder.MarkOrderPaidResp{
		Orders: updatedOrders,
	}

	// 返回响应
	return resp, nil
}
