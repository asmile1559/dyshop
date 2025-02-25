// charge_test.go
package service

import (
	"context"
	"errors"
	"math/rand"
	"reflect"
	"testing"
	"time"

	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	"github.com/asmile1559/dyshop/app/payment/biz/dal"
	"github.com/asmile1559/dyshop/app/payment/biz/model"
	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB 初始化一个内存 SQLite 数据库用于测试
func setupTestDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("打开内存数据库失败: %v", err)
	}
	// 自动迁移 PaymentRecord 模型
	err = db.AutoMigrate(&model.PaymentRecord{})
	if err != nil {
		t.Fatalf("迁移数据库失败: %v", err)
	}
	dal.DB = db
}

func TestChargeService_Run(t *testing.T) {
	// 子测试 1：正常支付流程
	t.Run("valid request", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		validReq := &pbpayment.ChargeReq{
			Amount:  100.0,
			OrderId: "order-123",
			UserId:  1,
			CreditCard: &pbpayment.CreditCardInfo{
				CreditCardNumber:          "4111111111111111",
				CreditCardCvv:             123,
				CreditCardExpirationYear:  int32(time.Now().Year() + 1),
				CreditCardExpirationMonth: int32(time.Now().Month()),
			},
		}

		// 固定随机返回值，确保支付成功（0.5 > 0.1）
		patch := monkey.Patch(rand.Float32, func() float32 {
			return 0.5
		})
		defer patch.Unpatch()

		// 固定生成的交易号
		patch2 := monkey.Patch(generateTransactionID, func() string {
			return "TXN_FIXED"
		})
		defer patch2.Unpatch()

		resp, err := service.Run(validReq)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "TXN_FIXED", resp.TransactionId)

		// 验证支付流水记录已插入数据库
		var record model.PaymentRecord
		result := dal.DB.First(&record, "order_id = ?", "order-123")
		assert.NoError(t, result.Error)
		assert.Equal(t, "TXN_FIXED", record.TransactionID)
		assert.Equal(t, validReq.OrderId, record.OrderID)
		assert.Equal(t, validReq.UserId, record.UserID)
		assert.Equal(t, validReq.Amount, record.Amount)
	})

	// 子测试 2：请求为 nil
	t.Run("nil request", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		resp, err := service.Run(nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "请求不能为空")
		assert.Nil(t, resp)
	})

	// 子测试 3：支付金额不合法
	t.Run("invalid amount", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		req := &pbpayment.ChargeReq{
			Amount:  0,
			OrderId: "order-123",
			UserId:  1,
			CreditCard: &pbpayment.CreditCardInfo{
				CreditCardNumber:          "4111111111111111",
				CreditCardCvv:             123,
				CreditCardExpirationYear:  int32(time.Now().Year() + 1),
				CreditCardExpirationMonth: int32(time.Now().Month()),
			},
		}
		resp, err := service.Run(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "支付金额必须大于0")
		assert.Nil(t, resp)
	})

	// 子测试 4：订单号为空
	t.Run("missing order id", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		req := &pbpayment.ChargeReq{
			Amount: 100,
			// OrderId 为空
			OrderId: "",
			UserId:  1,
			CreditCard: &pbpayment.CreditCardInfo{
				CreditCardNumber:          "4111111111111111",
				CreditCardCvv:             123,
				CreditCardExpirationYear:  int32(time.Now().Year() + 1),
				CreditCardExpirationMonth: int32(time.Now().Month()),
			},
		}
		resp, err := service.Run(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "订单号不能为空")
		assert.Nil(t, resp)
	})

	// 子测试 5：用户ID为0
	t.Run("missing user id", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		req := &pbpayment.ChargeReq{
			Amount:  100,
			OrderId: "order-123",
			UserId:  0,
			CreditCard: &pbpayment.CreditCardInfo{
				CreditCardNumber:          "4111111111111111",
				CreditCardCvv:             123,
				CreditCardExpirationYear:  int32(time.Now().Year() + 1),
				CreditCardExpirationMonth: int32(time.Now().Month()),
			},
		}
		resp, err := service.Run(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "用户ID不能为空")
		assert.Nil(t, resp)
	})

	// 子测试 6：信用卡信息为空
	t.Run("missing credit card info", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		req := &pbpayment.ChargeReq{
			Amount:     100,
			OrderId:    "order-123",
			UserId:     1,
			CreditCard: nil,
		}
		resp, err := service.Run(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "信用卡信息不能为空")
		assert.Nil(t, resp)
	})

	// 子测试 7：无效的信用卡号码（位数不足）
	t.Run("invalid credit card number", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		req := &pbpayment.ChargeReq{
			Amount:  100,
			OrderId: "order-123",
			UserId:  1,
			CreditCard: &pbpayment.CreditCardInfo{
				// 信用卡号码位数不足
				CreditCardNumber: "123456789",
				CreditCardCvv:    123,
				CreditCardExpirationYear:  int32(time.Now().Year() + 1),
				CreditCardExpirationMonth: int32(time.Now().Month()),
			},
		}
		resp, err := service.Run(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的信用卡号码")
		assert.Nil(t, resp)
	})

	// 子测试 8：无效的信用卡 CVV（位数不足）
	t.Run("invalid credit card CVV", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		req := &pbpayment.ChargeReq{
			Amount:  100,
			OrderId: "order-123",
			UserId:  1,
			CreditCard: &pbpayment.CreditCardInfo{
				CreditCardNumber:          "4111111111111111",
				// CVV 位数不足（只有2位）
				CreditCardCvv:             12,
				CreditCardExpirationYear:  int32(time.Now().Year() + 1),
				CreditCardExpirationMonth: int32(time.Now().Month()),
			},
		}
		resp, err := service.Run(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "无效的信用卡CVV")
		assert.Nil(t, resp)
	})

	// 子测试 9：信用卡已过期
	t.Run("expired credit card", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		req := &pbpayment.ChargeReq{
			Amount:  100,
			OrderId: "order-123",
			UserId:  1,
			CreditCard: &pbpayment.CreditCardInfo{
				CreditCardNumber:          "4111111111111111",
				CreditCardCvv:             123,
				// 使用过去的年份表示已过期
				CreditCardExpirationYear:  int32(time.Now().Year() - 1),
				CreditCardExpirationMonth: 12,
			},
		}
		resp, err := service.Run(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "信用卡已过期")
		assert.Nil(t, resp)
	})

	// 子测试 10：支付平台处理失败（模拟支付失败）
	t.Run("payment gateway failure", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		validReq := &pbpayment.ChargeReq{
			Amount:  100.0,
			OrderId: "order-123",
			UserId:  1,
			CreditCard: &pbpayment.CreditCardInfo{
				CreditCardNumber:          "4111111111111111",
				CreditCardCvv:             123,
				CreditCardExpirationYear:  int32(time.Now().Year() + 1),
				CreditCardExpirationMonth: int32(time.Now().Month()),
			},
		}

		// 固定随机返回值，模拟失败（0.05 < 0.1）
		patch := monkey.Patch(rand.Float32, func() float32 {
			return 0.05
		})
		defer patch.Unpatch()

		resp, err := service.Run(validReq)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "支付平台处理失败")
		assert.Nil(t, resp)
	})

	// 子测试 11：模拟数据库写入失败
	t.Run("db insertion failure", func(t *testing.T) {
		setupTestDB(t)
		service := NewChargeService(context.Background())
		validReq := &pbpayment.ChargeReq{
			Amount:  100.0,
			OrderId: "order-123",
			UserId:  1,
			CreditCard: &pbpayment.CreditCardInfo{
				CreditCardNumber:          "4111111111111111",
				CreditCardCvv:             123,
				CreditCardExpirationYear:  int32(time.Now().Year() + 1),
				CreditCardExpirationMonth: int32(time.Now().Month()),
			},
		}

		// 固定支付成功
		patch := monkey.Patch(rand.Float32, func() float32 {
			return 0.5
		})
		defer patch.Unpatch()

		// 模拟生成固定交易号
		patch2 := monkey.Patch(generateTransactionID, func() string {
			return "TXN_FIXED"
		})
		defer patch2.Unpatch()

		// Patch dal.DB.Create 方法模拟写入错误
		originalDB := dal.DB
		patchCreate := monkey.PatchInstanceMethod(reflect.TypeOf(originalDB), "Create", func(db *gorm.DB, value interface{}) *gorm.DB {
			return &gorm.DB{Error: errors.New("db insertion error")}
		})
		defer patchCreate.Unpatch()

		resp, err := service.Run(validReq)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "记录支付流水失败")
		assert.Nil(t, resp)
	})
}