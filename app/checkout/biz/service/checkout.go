package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	payment "github.com/asmile1559/dyshop/pb/backend/payment" 
	"github.com/google/uuid"
)

// validateEmail 使用简单的正则表达式校验 Email 格式
func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// validateAddress 校验地址信息是否完整
func validateAddress(addr *pbcheckout.Address) error {
	if addr == nil {
		return errors.New("address is nil")
	}
	if addr.StreetAddress == "" || addr.City == "" || addr.State == "" ||
		addr.Country == "" || addr.ZipCode == "" {
		return errors.New("incomplete address fields")
	}
	return nil
}

// simulatePayment 模拟支付处理，实际场景中应调用支付网关
func simulatePayment(creditCard *payment.CreditCardInfo, amount float64) (string, error) {
	// 检查信用卡信息是否存在
	if creditCard == nil {
		return "", errors.New("credit card information is missing")
	}
	// 此处可以添加更多校验逻辑，如校验卡号格式、有效期等

	// 模拟支付处理时间
	time.Sleep(500 * time.Millisecond)

	// 模拟支付成功，返回一个新的交易 ID
	transactionID := uuid.New().String()
	return transactionID, nil
}

type CheckoutService struct {
	ctx context.Context
}

func NewCheckoutService(c context.Context) *CheckoutService {
	return &CheckoutService{ctx: c}
}

func (s *CheckoutService) Run(req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {
	// 1. 校验用户信息
	if req.UserId == 0 {
		return nil, errors.New("invalid user id")
	}
	if req.Email == "" || !validateEmail(req.Email) {
		return nil, errors.New("invalid email format")
	}

	// 2. 校验地址信息
	if err := validateAddress(req.Address); err != nil {
		return nil, fmt.Errorf("address validation failed: %v", err)
	}

	// 3. 模拟支付处理（此处假设订单金额为 100.00）
	amount := 100.00
	transactionID, err := simulatePayment(req.CreditCard, amount)
	if err != nil {
		return nil, fmt.Errorf("payment processing failed: %v", err)
	}

	// 4. 生成订单号（使用 UUID 生成）
	orderID := uuid.New().String()

	// 5. (可选) 将订单信息保存到数据库，此处略过

	// 6. 返回订单结算响应
	return &pbcheckout.CheckoutResp{
		OrderId:       orderID,
		TransactionId: transactionID,
	}, nil
}
