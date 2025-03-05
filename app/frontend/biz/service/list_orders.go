package service

import (
	"context"
	"errors"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/asmile1559/dyshop/pb/frontend/order_page"
	"github.com/gin-gonic/gin"
)

type ListOrdersService struct {
	ctx context.Context
}

func NewListOrdersService(c context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: c}
}

func (s *ListOrdersService) Run(req *order_page.ListOrderReq) (gin.H, error) {
	id, ok := s.ctx.Value("user_id").(int64)
	if ok == false {
		return nil, errors.New("expect user id")
	}

	// 获取user信息
	// userClient, conn, err := rpcclient.GetUserClient()
	// if err != nil {
	// 	return nil, err
	// }
	// defer conn.Close()
	// userResp, err := userClient.GetUserInfo(s.ctx, &pbuser.GetUserInfoReq{
	// 	UserId: id,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	orderClient, conn, err := rpcclient.GetOrderClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	orderResp, err := orderClient.ListOrders(s.ctx, &pborder.ListOrderReq{
		UserId: id,
	})
	if err != nil {
		return nil, err
	}

	orders := orderResp.GetOrders()
	if len(orders) == 0 {
	}

	// _resp = gin.H{
	// 	"PageRouter": PageRouter,
	// 	"UserInfo": gin.H{
	// 		"Name": userResp.GetName(),
	// 	},

	// 	"AddressInfo": gin.H{ // 自维护
	// 		"Default": "1",
	// 		"Addresses": []gin.H{
	// 			{
	// 				"AddressId":   "1",
	// 				"Recipient":   "张三李四",
	// 				"Phone":       "12345678901",
	// 				"Province":    "中国",
	// 				"City":        "北京市",
	// 				"District":    "海淀区",
	// 				"Street":      "知春路",
	// 				"FullAddress": "北京北京市海淀区知春路甲48号抖音视界",
	// 			},
	// 		},
	// 	},
	// 	// cid -> cart rpc(GetCart()) -> cart.product_id -> product rpc(GetProduct) -> product info
	// 	"Products": []gin.H{
	// 		{
	// 			"ProductId":   "1",
	// 			"ProductImg":  "/static/src/product/bearcookie.webp",
	// 			"ProductName": "超级无敌好吃的小熊饼干",
	// 			"ProductSpec": gin.H{
	// 				"Name":  "无",
	// 				"Price": "18.80",
	// 			},
	// 			"Quantity": "2",
	// 			"Currency": "CNY",
	// 			"Postage":  "10.00",
	// 		},
	// 		{
	// 			"ItemId":      "2",
	// 			"productId":   "2",
	// 			"ProductImg":  "/static/src/product/bearsweet.webp",
	// 			"ProductName": "超级无敌好吃的小熊软糖值得品尝大力推荐",
	// 			"ProductSpec": gin.H{
	// 				"Name":  "9分软",
	// 				"Price": "20.99",
	// 			},
	// 			"Quantity": "1",
	// 			"Postage":  "0",
	// 		},
	// 	},
	// 	"OrderPrice":      "58.59",
	// 	"OrderPostage":    "10.00",
	// 	"OrderDiscount":   "0",
	// 	"OrderFinalPrice": "68.59",
	// }

	return gin.H{
		"orders": orderResp.GetOrders(),
	}, nil

}
