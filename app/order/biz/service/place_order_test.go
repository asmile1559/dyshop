package service

import (
	"context"
	"testing"

	"github.com/asmile1559/dyshop/app/order/biz/model"
	"github.com/asmile1559/dyshop/app/order/utils/db"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestPlaceOrderService_Run 测试 PlaceOrderService 的 Run 方法
func TestPlaceOrderService_Run(t *testing.T) {
	// 初始化日志记录器
	logrus.SetLevel(logrus.DebugLevel)

	// 设置 GORM 使用 SQLite 内存数据库进行测试
	var err error
	db.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database connection: %v", err)
	}

	// 自动迁移模型
	err = db.DB.AutoMigrate(&model.Order{}, &model.Address{}, &model.OrderItem{})
	if err != nil {
		t.Fatalf("failed to migrate models: %v", err)
	}

	// 创建一个模拟的 PlaceOrderReq 请求
	req := &pborder.PlaceOrderReq{
		UserId:       1,
		UserCurrency: "USD",
		Address: &pborder.Address{
			StreetAddress: "123 Main St",
			City:          "Anytown",
			State:         "CA",
			Country:       "USA",
			ZipCode:       "12345",
		},
		Email: "test@example.com",
		OrderItems: []*pborder.OrderItem{
			{
				Item: &pbcart.CartItem{
					ProductId: 101,
					Quantity:  2,
				},
				Cost: 19.99,
			},
			/*{
				Item: &pbcart.CartItem{
					ProductId: 102,
					Quantity:  1,
				},
				Cost: 29.99,
			},*/
		},
	}

	// 创建 PlaceOrderService 实例
	ctx := context.Background()
	service := NewPlaceOrderService(ctx)

	// 调用 Run 方法
	resp, err := service.Run(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// 验证响应
	expectedOrderID := resp.GetOrder().GetOrderId()
	if expectedOrderID == "" {
		t.Errorf("expected non-empty order ID, got empty")
	}

	// 查询数据库以验证订单是否已成功创建
	var createdOrder model.Order
	result := db.DB.Preload("Address").Preload("OrderItems").First(&createdOrder, expectedOrderID)
	if result.Error != nil {
		t.Errorf("expected order to be found in the database, got error: %v", result.Error)
	}

	// 验证订单字段
	if createdOrder.UserId != uint64(req.GetUserId()) {
		t.Errorf("expected UserId %d, got %d", req.GetUserId(), createdOrder.UserId)
	}
	if createdOrder.UserCurrency != req.GetUserCurrency() {
		t.Errorf("expected UserCurrency %s, got %s", req.GetUserCurrency(), createdOrder.UserCurrency)
	}
	if createdOrder.Email != req.GetEmail() {
		t.Errorf("expected Email %s, got %s", req.GetEmail(), createdOrder.Email)
	}
	if len(createdOrder.OrderItems) != len(req.GetOrderItems()) {
		t.Errorf("expected %d OrderItems, got %d", len(req.GetOrderItems()), len(createdOrder.OrderItems))
	}

	// 验证地址字段
	if createdOrder.Address.StreetAddress != req.GetAddress().GetStreetAddress() {
		t.Errorf("expected StreetAddress %s, got %s", req.GetAddress().GetStreetAddress(), createdOrder.Address.StreetAddress)
	}
	if createdOrder.Address.City != req.GetAddress().GetCity() {
		t.Errorf("expected City %s, got %s", req.GetAddress().GetCity(), createdOrder.Address.City)
	}
	if createdOrder.Address.State != req.GetAddress().GetState() {
		t.Errorf("expected State %s, got %s", req.GetAddress().GetState(), createdOrder.Address.State)
	}
	if createdOrder.Address.Country != req.GetAddress().GetCountry() {
		t.Errorf("expected Country %s, got %s", req.GetAddress().GetCountry(), createdOrder.Address.Country)
	}
	if createdOrder.Address.ZipCode != req.GetAddress().GetZipCode() {
		t.Errorf("expected ZipCode %s, got %s", req.GetAddress().GetZipCode(), createdOrder.Address.ZipCode)
	}

	// 验证订单项字段
	for i, item := range createdOrder.OrderItems {
		expectedItem := req.GetOrderItems()[i].GetItem()
		if item.ProductId != uint64(expectedItem.GetProductId()) {
			t.Errorf("expected ProductId %d for item %d, got %d", expectedItem.GetProductId(), i, item.ProductId)
		}
		if item.Quantity != int(expectedItem.GetQuantity()) {
			t.Errorf("expected Quantity %d for item %d, got %d", expectedItem.GetQuantity(), i, item.Quantity)
		}
		if item.Cost != float64(req.GetOrderItems()[i].GetCost()) {
			t.Errorf("expected Cost %.2f for item %d, got %.2f", req.GetOrderItems()[i].GetCost(), i, item.Cost)
		}
	}
}
