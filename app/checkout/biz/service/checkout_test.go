// service_test.go
package service_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	"github.com/asmile1559/dyshop/app/checkout/biz/model"
	"github.com/asmile1559/dyshop/app/checkout/biz/service"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	"bou.ke/monkey"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ----------------- Fake Payment 服务 -----------------

type fakePaymentServer struct {
	pbpayment.UnimplementedPaymentServiceServer
}

func (s *fakePaymentServer) Charge(ctx context.Context, req *pbpayment.ChargeReq) (*pbpayment.ChargeResp, error) {
	// 模拟信用卡号 "0000000000000000" 支付失败
	if req.CreditCard.CreditCardNumber == "0000000000000000" {
		return nil, errors.New("模拟支付失败")
	}
	// 正常返回一个假交易ID
	return &pbpayment.ChargeResp{
		TransactionId: "fake-tx-123",
	}, nil
}

// ----------------- Fake Order 服务 -----------------

type fakeOrderServer struct {
	pborder.UnimplementedOrderServiceServer
}

func (s *fakeOrderServer) ListOrder(ctx context.Context, req *pborder.ListOrderReq) (*pborder.ListOrderResp, error) {
	// 模拟返回一笔订单（注意：Cost 为 float32 类型）
	fakeOrder := &pborder.Order{
		OrderId: "order-123",
		OrderItems: []*pborder.OrderItem{
			{Cost: 100}, // 订单项金额 100
		},
	}
	return &pborder.ListOrderResp{
		Orders: []*pborder.Order{fakeOrder},
	}, nil
}

func startFakeGRPCServer(register func(*grpc.Server), t *testing.T) (net.Listener, *grpc.Server) {
	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	register(gs)
	go func() {
		if err := gs.Serve(lis); err != nil {
			t.Logf("GRPC server stopped: %v", err)
		}
	}()
	return lis, gs
}

func setupTestDB(t *testing.T) {
	// 使用内存 SQLite 数据库作为测试数据库
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite in-memory database: %v", err)
	}
	// 自动迁移 OrderRecord 模型
	err = db.AutoMigrate(&model.OrderRecord{})
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	// 将 dal.DB 指向测试数据库
	dal.DB = db
}

func TestCheckoutService_Run(t *testing.T) {
	// 初始化测试数据库
	setupTestDB(t)

	// 启动假 Payment 服务
	paymentLis, paymentServer := startFakeGRPCServer(func(gs *grpc.Server) {
		pbpayment.RegisterPaymentServiceServer(gs, &fakePaymentServer{})
	}, t)
	defer paymentServer.Stop()
	paymentPort := paymentLis.Addr().(*net.TCPAddr).Port

	// 启动假 Order 服务
	orderLis, orderServer := startFakeGRPCServer(func(gs *grpc.Server) {
		pborder.RegisterOrderServiceServer(gs, &fakeOrderServer{})
	}, t)
	defer orderServer.Stop()
	orderPort := orderLis.Addr().(*net.TCPAddr).Port

	// 利用 monkey patch 替换 InitViper 函数，根据配置文件路径返回期望的端口
	patch := monkey.Patch(service.InitViper, func(configPath string) (*viper.Viper, error) {
		v := viper.New()
		// 根据路径判断返回 Payment 或 Order 服务的端口
		if strings.Contains(configPath, "payment") {
			v.Set("server.port", strconv.Itoa(paymentPort))
		} else if strings.Contains(configPath, "order") {
			v.Set("server.port", strconv.Itoa(orderPort))
		} else {
			return nil, fmt.Errorf("unknown config path: %s", configPath)
		}
		return v, nil
	})
	defer patch.Unpatch()

	// 创建结算服务实例
	svc := service.NewCheckoutService(context.Background())

	// 构造一个有效的结算请求
	validReq := &pbcheckout.CheckoutReq{
		UserId:    1,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john@example.com",
		Address: &pbcheckout.Address{
			StreetAddress: "123 Main St",
			City:          "Metropolis",
			ZipCode:       "12345",
		},
		CreditCard: &pbpayment.CreditCardInfo{
			CreditCardNumber:          "4111111111111111",
			CreditCardCvv:             123,
			CreditCardExpirationYear:  2026,
			CreditCardExpirationMonth: 12,
		},
	}

	// 子测试 1：正常流程
	t.Run("valid request", func(t *testing.T) {
		resp, err := svc.Run(validReq)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "order-123", resp.OrderId)
		assert.Equal(t, "fake-tx-123", resp.TransactionId)

		// 验证数据库中是否写入订单记录
		var record model.OrderRecord
		dbResult := dal.DB.First(&record, "order_id = ?", "order-123")
		assert.NoError(t, dbResult.Error)
		assert.EqualValues(t, validReq.UserId, record.UserID)
		assert.Equal(t, "fake-tx-123", record.TransactionID)
		// 订单金额：100（订单项只有一笔 100）
		assert.Equal(t, float64(100), record.Amount)
	})

	// 子测试 2：请求参数不合法（缺少必要字段）
	t.Run("invalid request - missing fields", func(t *testing.T) {
		invalidReq := &pbcheckout.CheckoutReq{
			// 缺少 UserId, Firstname, Lastname, Email 等必要信息
			Address: &pbcheckout.Address{
				StreetAddress: "123 Main St",
				City:          "Metropolis",
				ZipCode:       "12345",
			},
			CreditCard: &pbpayment.CreditCardInfo{
				CreditCardNumber: "4111111111111111",
				CreditCardCvv:    123,
			},
		}
		_, err := svc.Run(invalidReq)
		assert.Error(t, err)
		assert.Equal(t, "缺少必要字段", err.Error())
	})

	// 子测试 3：模拟支付失败（信用卡号特殊处理）
	t.Run("payment failure", func(t *testing.T) {
		req := &pbcheckout.CheckoutReq{
			UserId:    2,
			Firstname: "Alice",
			Lastname:  "Smith",
			Email:     "alice@example.com",
			Address: &pbcheckout.Address{
				StreetAddress: "456 Another St",
				City:          "Gotham",
				ZipCode:       "54321",
			},
			CreditCard: &pbpayment.CreditCardInfo{
				// 特殊信用卡号触发支付失败
				CreditCardNumber:          "0000000000000000",
				CreditCardCvv:             999,
				CreditCardExpirationYear:  2025,
				CreditCardExpirationMonth: 10,
			},
		}
		resp, err := svc.Run(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "支付失败", err.Error())
	})
	
	// 给足够时间让异步数据库写入完成
	time.Sleep(100 * time.Millisecond)
}