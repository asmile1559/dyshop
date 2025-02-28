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
	//"reflect"
	"time"
)

// TestListOrdersService_Run 测试 ListOrdersService 的 Run 方法
func TestListOrdersService_Run(t *testing.T) {
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

	// 创建一些示例订单数据
	now := time.Now()
	testOrders := []model.Order{
		{
			UserId:       1,
			UserCurrency: "USD",
			Address: model.Address{
				StreetAddress: "123 Main St",
				City:          "Anytown",
				State:         "CA",
				Country:       "USA",
				ZipCode:       "12345",
			},
			Email:     "test@example.com",
			CreatedAt: now,
			OrderItems: []model.OrderItem{
				{ProductId: 101, Quantity: 2, Cost: 19.99},
			},
		},
		{
			UserId:       2,
			UserCurrency: "EUR",
			Address: model.Address{
				StreetAddress: "456 Side St",
				City:          "Othertown",
				State:         "TX",
				Country:       "USA",
				ZipCode:       "67890",
			},
			Email:     "another@example.com",
			CreatedAt: now,
			OrderItems: []model.OrderItem{
				{ProductId: 102, Quantity: 1, Cost: 29.99},
			},
		},
	}

	// 插入测试数据到数据库
	for _, order := range testOrders {
		if err := db.DB.Create(&order).Error; err != nil {
			t.Fatalf("failed to create test order: %v", err)
		}
	}

	// 查询数据库以验证订单是否已成功创建
	var fetchedOrders []model.Order
	if err := db.DB.Preload("Address").Preload("OrderItems").Find(&fetchedOrders).Error; err != nil {
		t.Fatalf("failed to fetch orders from database: %v", err)
	}

	if len(fetchedOrders) != len(testOrders) {
		t.Fatalf("expected %d orders in the database, got %d", len(testOrders), len(fetchedOrders))
	}

	// 创建一个模拟的 ListOrderReq 请求
	req := &pborder.ListOrderReq{
		UserId: 1,
	}

	// 创建 ListOrdersService 实例
	ctx := context.Background()
	service := NewListOrdersService(ctx)

	// 调用 Run 方法
	resp, err := service.Run(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// 验证响应
	expectedOrders := []*pborder.Order{
		{
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

	if len(resp.Orders) != len(expectedOrders) {
		t.Errorf("expected %d orders, got %d", len(expectedOrders), len(resp.Orders))
		return
	}

	for i := range expectedOrders {
		if !compareOrders(resp.Orders[i], expectedOrders[i]) {
			t.Errorf("expected orders do not match, got %+v", resp.Orders[i])
		}
	}
}
