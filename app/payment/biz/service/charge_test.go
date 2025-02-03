package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	"github.com/asmile1559/dyshop/app/payment/biz/dal"
	"github.com/asmile1559/dyshop/app/payment/biz/model"
	"github.com/asmile1559/dyshop/app/payment/mocks"
)

func TestChargeRun(t *testing.T) {
	// 初始化上下文
	ctx := context.Background()

	// 创建 gomock 控制器
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建模拟的 DB 对象
	mockDB := mocks.NewMockDB(ctrl)

	// 替换全局的 dal.DB 为模拟对象
	oldDB := dal.DB
	dal.DB = mockDB
	defer func() { dal.DB = oldDB }()

	// 初始化 ChargeService
	service := NewChargeService(ctx)

	// 定义测试用例
	tests := []struct {
		name        string
		req         *pbpayment.ChargeReq
		mockDBSetup func()
		wantResp    *pbpayment.ChargeResp
		wantErr     bool
	}{
		{
			name: "成功支付",
			req: &pbpayment.ChargeReq{
				Amount:  100,
				OrderId: "order123",
				UserId:  1,
				CreditCard: &pbpayment.CreditCard{
					CreditCardNumber:     "4111111111111111",
					CreditCardExpirationMonth: 12,
					CreditCardExpirationYear:  2025,
					CreditCardCvv:             123,
				},
			},
			mockDBSetup: func() {
				// 模拟数据库插入成功
				mockDB.EXPECT().Create(gomock.Any()).Return(nil)
			},
			wantResp: &pbpayment.ChargeResp{
				TransactionId: gomock.Any().String(),
			},
			wantErr: false,
		},
		{
			name: "支付金额无效",
			req: &pbpayment.ChargeReq{
				Amount:  0,
				OrderId: "order123",
				UserId:  1,
				CreditCard: &pbpayment.CreditCard{
					CreditCardNumber:     "4111111111111111",
					CreditCardExpirationMonth: 12,
					CreditCardExpirationYear:  2025,
					CreditCardCvv:             123,
				},
			},
			mockDBSetup: func() {
				// 不需要模拟数据库调用，因为会在校验阶段失败
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "支付平台失败",
			req: &pbpayment.ChargeReq{
				Amount:  100,
				OrderId: "order123",
				UserId:  1,
				CreditCard: &pbpayment.CreditCard{
					CreditCardNumber:     "4111111111111111",
					CreditCardExpirationMonth: 12,
					CreditCardExpirationYear:  2025,
					CreditCardCvv:             123,
				},
			},
			mockDBSetup: func() {
				// 不需要模拟数据库调用，因为支付平台会失败
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "数据库插入失败",
			req: &pbpayment.ChargeReq{
				Amount:  100,
				OrderId: "order123",
				UserId:  1,
				CreditCard: &pbpayment.CreditCard{
					CreditCardNumber:     "4111111111111111",
					CreditCardExpirationMonth: 12,
					CreditCardExpirationYear:  2025,
					CreditCardCvv:             123,
				},
			},
			mockDBSetup: func() {
				// 模拟数据库插入失败
				mockDB.EXPECT().Create(gomock.Any()).Return(errors.New("数据库错误"))
			},
			wantResp: nil,
			wantErr:  true,
		},
	}

	// 运行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 设置模拟对象的行为
			if tt.mockDBSetup != nil {
				tt.mockDBSetup()
			}

			// 调用 Run 方法
			resp, err := service.Run(tt.req)

			// 检查错误
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// 检查返回值
			if tt.wantResp != nil {
				assert.NotEmpty(t, resp.TransactionId)
			} else {
				assert.Nil(t, resp)
			}
		})
	}
}