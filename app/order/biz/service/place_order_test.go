package service

import (
	"context"
	"testing"

	"github.com/asmile1559/dyshop/app/order/biz/model"
	"github.com/asmile1559/dyshop/app/order/utils/db"
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
	err = db.DB.AutoMigrate(&model.Order{}, &model.Address{}, &model.Product{}, &model.ProductSpec{})
	if err != nil {
		t.Fatalf("failed to migrate models: %v", err)
	}

	// 创建一个模拟的 PlaceOrderReq 请求
	req := &pborder.PlaceOrderReq{
		UserId:       1,
		UserCurrency: "USD",
		Address: &pborder.Address{
			Recipient:   "John Doe",
			Phone:       "+123456789",
			Province:    "CA",
			City:        "Anytown",
			District:    "",
			Street:      "123 Main St",
			FullAddress: "123 Main St, Anytown, CA",
		},
		Email: "test@example.com",
		Products: []*pborder.Product{
			{
				ProductId:   101,
				ProductImg:  "img1.jpg",
				ProductName: "Product 1",
				ProductSpec: &pborder.ProductSpec{
					Name:  "Size S",
					Price: 19.99,
				},
				Quantity: 2,
				Currency: "USD",
				Postage:  5.00,
			},
			{
				ProductId:   102,
				ProductImg:  "img2.jpg",
				ProductName: "Product 2",
				ProductSpec: &pborder.ProductSpec{
					Name:  "Size M",
					Price: 29.99,
				},
				Quantity: 1,
				Currency: "USD",
				Postage:  5.00,
			},
		},
		OrderPrice:      69.97,
		OrderPostage:    5.00,
		Discount:        0.00,
		OrderFinalPrice: 74.97,
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
	expectedOrderID := resp.GetOrderId()
	if expectedOrderID == "" {
		t.Errorf("expected non-empty order ID, got empty")
	}

	// 查询数据库以验证订单是否已成功创建
	var createdOrder model.Order
	result := db.DB.Preload("Products").First(&createdOrder, "order_id = ?", expectedOrderID)
	if result.Error != nil {
		t.Errorf("expected order to be found in the database, got error: %v", result.Error)
	}

	// 验证订单字段
	if createdOrder.UserId != uint32(req.GetUserId()) {
		t.Errorf("expected UserId %d, got %d", req.GetUserId(), createdOrder.UserId)
	}
	if createdOrder.UserCurrency != req.GetUserCurrency() {
		t.Errorf("expected UserCurrency %s, got %s", req.GetUserCurrency(), createdOrder.UserCurrency)
	}
	if createdOrder.Email != req.GetEmail() {
		t.Errorf("expected Email %s, got %s", req.GetEmail(), createdOrder.Email)
	}
	if len(createdOrder.Products) != len(req.GetProducts()) {
		t.Errorf("expected %d Products, got %d", len(req.GetProducts()), len(createdOrder.Products))
	}
	if createdOrder.OrderPrice != req.GetOrderPrice() {
		t.Errorf("expected OrderPrice %.2f, got %.2f", req.GetOrderPrice(), createdOrder.OrderPrice)
	}
	if createdOrder.OrderPostage != req.GetOrderPostage() {
		t.Errorf("expected OrderPostage %.2f, got %.2f", req.GetOrderPostage(), createdOrder.OrderPostage)
	}
	if createdOrder.OrderDiscount != req.GetDiscount() {
		t.Errorf("expected Discount %.2f, got %.2f", req.GetDiscount(), createdOrder.OrderDiscount)
	}
	if createdOrder.OrderFinalPrice != req.GetOrderFinalPrice() {
		t.Errorf("expected FinalPrice %.2f, got %.2f", req.GetOrderFinalPrice(), createdOrder.OrderFinalPrice)
	}

	// 验证地址字段
	address := req.GetAddress()
	if createdOrder.Address.Recipient != address.GetRecipient() {
		t.Errorf("expected Recipient %s, got %s", address.GetRecipient(), createdOrder.Address.Recipient)
	}
	if createdOrder.Address.Phone != address.GetPhone() {
		t.Errorf("expected Phone %s, got %s", address.GetPhone(), createdOrder.Address.Phone)
	}
	if createdOrder.Address.Province != address.GetProvince() {
		t.Errorf("expected Province %s, got %s", address.GetProvince(), createdOrder.Address.Province)
	}
	if createdOrder.Address.City != address.GetCity() {
		t.Errorf("expected City %s, got %s", address.GetCity(), createdOrder.Address.City)
	}
	if createdOrder.Address.Street != address.GetStreet() {
		t.Errorf("expected Street %s, got %s", address.GetStreet(), createdOrder.Address.Street)
	}
	if createdOrder.Address.FullAddress != address.GetFullAddress() {
		t.Errorf("expected FullAddress %s, got %s", address.GetFullAddress(), createdOrder.Address.FullAddress)
	}

	// 验证产品字段
	for i, item := range createdOrder.Products {
		expectedItem := req.GetProducts()[i]
		if item.ProductID != uint64(expectedItem.GetProductId()) {
			t.Errorf("expected ProductId %d for item %d, got %d", expectedItem.GetProductId(), i, item.ProductID)
		}
		if item.ProductName != expectedItem.GetProductName() {
			t.Errorf("expected ProductName %s for item %d, got %s", expectedItem.GetProductName(), i, item.ProductName)
		}
		if item.ProductSpec.Name != expectedItem.GetProductSpec().GetName() {
			t.Errorf("expected ProductSpec Name %s for item %d, got %s", expectedItem.GetProductSpec().GetName(), i, item.ProductSpec.Name)
		}
		if item.Quantity != int(expectedItem.GetQuantity()) {
			t.Errorf("expected Quantity %d for item %d, got %d", expectedItem.GetQuantity(), i, item.Quantity)
		}
		if item.Currency != expectedItem.GetCurrency() {
			t.Errorf("expected Currency %s for item %d, got %s", expectedItem.GetCurrency(), i, item.Currency)
		}
		if item.Postage != expectedItem.GetPostage() {
			t.Errorf("expected Postage %.2f for item %d, got %.2f", expectedItem.GetPostage(), i, item.Postage)
		}
	}
}
