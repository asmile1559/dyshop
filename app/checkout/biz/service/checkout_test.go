package service

import (
	"context"
	"errors"
	"regexp"
	"testing"

	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	payment "github.com/asmile1559/dyshop/pb/backend/payment"
)

func TestCheckoutRun(t *testing.T) {
	// 构造一个完整有效的请求，用于成功案例
	validReq := &pbcheckout.CheckoutReq{
		UserId:    1,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "john.doe@example.com",
		Address: &pbcheckout.Address{
			StreetAddress: "123 Main St",
			City:          "CityName",
			State:         "StateName",
			Country:       "CountryName",
			ZipCode:       "12345",
		},
		CreditCard: &payment.CreditCardInfo{
			CardNumber: "4111111111111111",
			ExpMonth:   12,
			ExpYear:    2030,
			Cvv:        "123",
		},
	}

	// 定义多个测试用例
	tests := []struct {
		name        string
		req         *pbcheckout.CheckoutReq
		expectError bool
		errContains string // 部分错误信息匹配
	}{
		{
			name:        "valid request",
			req:         validReq,
			expectError: false,
		},
		{
			name: "invalid user id",
			req: func() *pbcheckout.CheckoutReq {
				req := *validReq
				req.UserId = 0
				return &req
			}(),
			expectError: true,
			errContains: "invalid user id",
		},
		{
			name: "invalid email format",
			req: func() *pbcheckout.CheckoutReq {
				req := *validReq
				req.Email = "invalid-email"
				return &req
			}(),
			expectError: true,
			errContains: "invalid email format",
		},
		{
			name: "incomplete address",
			req: func() *pbcheckout.CheckoutReq {
				req := *validReq
				req.Address = &pbcheckout.Address{
					StreetAddress: "",
					City:          "CityName",
					State:         "StateName",
					Country:       "CountryName",
					ZipCode:       "12345",
				}
				return &req
			}(),
			expectError: true,
			errContains: "incomplete address fields",
		},
		{
			name: "missing credit card info",
			req: func() *pbcheckout.CheckoutReq {
				req := *validReq
				req.CreditCard = nil
				return &req
			}(),
			expectError: true,
			errContains: "credit card information is missing",
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewCheckoutService(ctx)
			resp, err := service.Run(tt.req)

			// 检查是否符合预期是否有错误返回
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got nil")
				} else if !errors.Is(err, errors.New(tt.errContains)) &&
					!regexp.MustCompile(tt.errContains).MatchString(err.Error()) {
					t.Errorf("expected error to contain %q, got %v", tt.errContains, err)
				}
				// 对于错误情况，不需要检查响应
				return
			}

			// 对于成功情况，不应该有错误，且返回的订单和交易号应不为空
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if resp.OrderId == "" {
				t.Errorf("expected non-empty OrderId")
			}
			if resp.TransactionId == "" {
				t.Errorf("expected non-empty TransactionId")
			}
		})
	}
}
