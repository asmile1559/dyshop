package service

import (
	"context"
	"fmt"

	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type GetOrderWithItemsService struct {
	ctx context.Context
}

// NewGetOrderWithItemsService 创建订单查询服务实例
func NewGetOrderWithItemsService(ctx context.Context) *GetOrderWithItemsService {
	return &GetOrderWithItemsService{ctx: ctx}
}

// Run 方法返回 gin.H 结构
func (s *GetOrderWithItemsService) Run(orderId string) (gin.H, error) {
	// 调用 gRPC 查询订单
	checkoutClient, conn, err := rpcclient.GetCheckoutClient()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	resp, err := checkoutClient.GetOrderWithItems(s.ctx, &pbcheckout.GetOrderReq{
		OrderId: orderId,
	})
	if err != nil {
		st, _ := status.FromError(err)
		return nil, st.Err()
	}

	// 组装 gin.H 作为返回
	return gin.H{
		"OrderId":       resp.Order.OrderId,
		"UserId":        resp.Order.UserId,
		"TransactionId": resp.Order.TransactionId,
		"Address": gin.H{
			"Recipient":   resp.Order.Recipient,
			"Phone":       resp.Order.Phone,
			"Province":    resp.Order.Province,
			"City":        resp.Order.City,
			"District":    resp.Order.District,
			"Street":      resp.Order.Street,
			"FullAddress": resp.Order.FullAddress,
		},
		"Products":        buildOrderItems(resp.Items),
		"OrderQuantity":   resp.Order.TotalQuantity,
		"OrderPostage":    resp.Order.Postage,
		"OrderPrice":      resp.Order.TotalPrice,
		"OrderFinalPrice": resp.Order.FinalPrice,
	}, nil
}

// buildOrderItems 转换订单商品信息
func buildOrderItems(items []*pbcheckout.OrderItem) []gin.H {
	var orderItems []gin.H
	for _, item := range items {
		orderItems = append(orderItems, gin.H{
			"ProductId":   item.ProductId,
			"ProductImg":  item.ProductImg,
			"ProductName": item.ProductName,
			"ProductSpec": gin.H{
				"Name":  item.SpecName,
				"Price": fmt.Sprintf("%.2f", item.SpecPrice),
			},
			"Quantity": fmt.Sprintf("%d", item.Quantity),
			"Currency": item.Currency,
			"Postage":  item.Postage,
		})
	}
	return orderItems
}
