package service

import (
	"context"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"testing"

	"github.com/asmile1559/dyshop/app/order/biz/model"
	"github.com/asmile1559/dyshop/app/order/utils/db"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//"reflect"
	"time"
)

// TestMarkOrderPaidService_Run 测试 MarkOrderPaidService 的 Run 方法
func TestMarkOrderPaidService_Run(t *testing.T) {
	// 初始化日志记录器
	logrus.SetLevel(logrus.DebugLevel)

	// 设置 GORM 使用 SQLite 内存数据库进行测试
	var err error
	db.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database connection: %v", err)
	}

	// 自动迁移模型
	err = db.DB.AutoMigrate(&model.PrePaidOrder{}, &model.PrePaidAddress{}, &model.PrePaidOrderItem{}, &model.Order{}, &model.Address{}, &model.OrderItem{})
	if err != nil {
		t.Fatalf("failed to migrate models: %v", err)
	}

	// 创建一些示例预支付订单数据
	now := time.Now()
	testPrePaidAddress := model.PrePaidAddress{
		StreetAddress: "123 Main St",
		City:          "Anytown",
		State:         "CA",
		Country:       "USA",
		ZipCode:       "12345",
	}

	testPrePaidOrderItems := []model.PrePaidOrderItem{
		{ProductId: 101, Quantity: 2, Cost: 19.99},
	}

	testPrePaidOrder := model.PrePaidOrder{
		ID:           1,
		UserId:       1,
		UserCurrency: "USD",
		Address:      testPrePaidAddress,
		Email:        "test@example.com",
		CreatedAt:    now,
		OrderItems:   testPrePaidOrderItems,
	}

	// 插入测试数据到数据库
	if err := db.DB.Create(&testPrePaidOrder).Error; err != nil {
		t.Fatalf("failed to create test prepaid order: %v", err)
	}

	// 创建一个模拟的 MarkOrderPaidReq 请求
	req := &pborder.MarkOrderPaidReq{
		OrderId: "1",
	}

	// 创建 MarkOrderPaidService 实例
	ctx := context.Background()
	service := NewMarkOrderPaidService(ctx)

	// 调用 Run 方法
	resp, err := service.Run(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// 验证响应
	expectedOrder := &pborder.Order{
		OrderId:      "1",
		UserId:       1,
		UserCurrency: "USD",
		Address: &pborder.Address{
			StreetAddress: "123 Main St",
			City:          "Anytown",
			State:         "CA",
			Country:       "USA",
			ZipCode:       "12345",
		},
		Email:     "test@example.com",
		CreatedAt: int32(now.Unix()),
		OrderItems: []*pborder.OrderItem{
			{
				Item: &pbcart.CartItem{
					ProductId: 101,
					Quantity:  2,
				},
				Cost: float32(19.99),
			},
		},
	}

	// 自定义比较函数来验证订单的关键字段
	compareOrders := func(actual, expected *pborder.Order) bool {
		if actual.OrderId != expected.OrderId {
			t.Errorf("OrderId mismatch: got %s, want %s", actual.OrderId, expected.OrderId)
			return false
		}
		if actual.UserId != expected.UserId {
			t.Errorf("UserId mismatch: got %d, want %d", actual.UserId, expected.UserId)
			return false
		}
		if actual.UserCurrency != expected.UserCurrency {
			t.Errorf("UserCurrency mismatch: got %s, want %s", actual.UserCurrency, expected.UserCurrency)
			return false
		}
		if actual.Email != expected.Email {
			t.Errorf("Email mismatch: got %s, want %s", actual.Email, expected.Email)
			return false
		}
		if actual.Address.StreetAddress != expected.Address.StreetAddress {
			t.Errorf("StreetAddress mismatch: got %s, want %s", actual.Address.StreetAddress, expected.Address.StreetAddress)
			return false
		}
		if actual.Address.City != expected.Address.City {
			t.Errorf("City mismatch: got %s, want %s", actual.Address.City, expected.Address.City)
			return false
		}
		if actual.Address.State != expected.Address.State {
			t.Errorf("State mismatch: got %s, want %s", actual.Address.State, expected.Address.State)
			return false
		}
		if actual.Address.Country != expected.Address.Country {
			t.Errorf("Country mismatch: got %s, want %s", actual.Address.Country, expected.Address.Country)
			return false
		}
		if actual.Address.ZipCode != expected.Address.ZipCode {
			t.Errorf("ZipCode mismatch: got %s, want %s", actual.Address.ZipCode, expected.Address.ZipCode)
			return false
		}
		if len(actual.OrderItems) != len(expected.OrderItems) {
			t.Errorf("OrderItems length mismatch: got %d, want %d", len(actual.OrderItems), len(expected.OrderItems))
			return false
		}
		for i, item := range actual.OrderItems {
			if item.Item.ProductId != expected.OrderItems[i].Item.ProductId {
				t.Errorf("OrderItem ProductId mismatch at index %d: got %d, want %d", i, item.Item.ProductId, expected.OrderItems[i].Item.ProductId)
				return false
			}
			if item.Item.Quantity != expected.OrderItems[i].Item.Quantity {
				t.Errorf("OrderItem Quantity mismatch at index %d: got %d, want %d", i, item.Item.Quantity, expected.OrderItems[i].Item.Quantity)
				return false
			}
			if item.Cost != expected.OrderItems[i].Cost {
				t.Errorf("OrderItem Cost mismatch at index %d: got %f, want %f", i, item.Cost, expected.OrderItems[i].Cost)
				return false
			}
		}
		if actual.CreatedAt != expected.CreatedAt {
			t.Errorf("CreatedAt mismatch: got %d, want %d", actual.CreatedAt, expected.CreatedAt)
			return false
		}
		return true
	}

	if len(resp.Orders) != 1 {
		t.Errorf("expected 1 order in response, got %d", len(resp.Orders))
		return
	}

	if !compareOrders(resp.Orders[0], expectedOrder) {
		t.Errorf("orders do not match")
	}

	// 验证预支付订单是否已被删除
	var fetchedPrePaidOrder model.PrePaidOrder
	if err := db.DB.First(&fetchedPrePaidOrder, req.GetOrderId()).Error; err == nil {
		t.Errorf("prepaid order was not deleted")
	} else if err != gorm.ErrRecordNotFound {
		t.Errorf("unexpected error while fetching deleted prepaid order: %v", err)
	}
}
