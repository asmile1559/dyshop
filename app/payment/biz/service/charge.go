package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	"github.com/asmile1559/dyshop/app/payment/biz/dal"
	"github.com/asmile1559/dyshop/app/payment/biz/model"
)

// ChargeService 处理支付请求的业务逻辑
type ChargeService struct {
	ctx context.Context
}

// NewChargeService 构造函数
func NewChargeService(c context.Context) *ChargeService {
	return &ChargeService{ctx: c}
}

// Run 处理支付请求并返回交易结果
func (s *ChargeService) Run(req *pbpayment.ChargeReq) (*pbpayment.ChargeResp, error) {
	// 1. 校验请求参数
	if err := s.validateRequest(req); err != nil {
		_ = s.recordPayment(req, "", "FAILED", fmt.Sprintf("请求参数错误: %v", err))
		return nil, err
	}

	// 2. 模拟调用第三方支付平台进行扣款
	transactionID, err := s.callPaymentGateway(req)
	if err != nil {
		// 记录失败的交易
		_ = s.recordPayment(req, "", "FAILED", fmt.Sprintf("支付失败: %v", err))
		return nil, err
	}

	// 3. 记录支付流水到数据库
	if err := s.recordPayment(req, transactionID, "SUCCESS", ""); err != nil {
		return nil, err
	}

	return &pbpayment.ChargeResp{TransactionId: transactionID}, nil
}

// validateRequest 对请求参数进行基本的校验
func (s *ChargeService) validateRequest(req *pbpayment.ChargeReq) error {
	if req == nil {
		return errors.New("请求不能为空")
	}
	if req.FinalPrice == "" {
		return errors.New("支付金额不能为空")
	}
	if _, err := strconv.ParseFloat(req.FinalPrice, 64); err != nil {
		return errors.New("支付金额格式不正确")
	}

	// 校验信用卡信息
	cc := req.CreditCard
	if cc == nil {
		return errors.New("信用卡信息不能为空")
	}
	// 校验卡号长度
	ccNumber := strings.TrimSpace(cc.CreditCardNumber)
	if len(ccNumber) < 13 || len(ccNumber) > 19 {
		return errors.New("无效的信用卡号码")
	}
	// 校验 CVV
	cvvStr := strconv.Itoa(int(cc.CreditCardCvv))
	if len(cvvStr) < 3 || len(cvvStr) > 4 {
		return errors.New("无效的信用卡CVV")
	}
	// 校验过期时间
	now := time.Now()
	expYear := int(cc.CreditCardExpirationYear)
	expMonth := int(cc.CreditCardExpirationMonth)
	if expYear < now.Year() || (expYear == now.Year() && expMonth < int(now.Month())) {
		return errors.New("信用卡已过期")
	}

	return nil
}

// callPaymentGateway 模拟调用第三方支付平台进行扣款
func (s *ChargeService) callPaymentGateway(req *pbpayment.ChargeReq) (string, error) {
	time.Sleep(500 * time.Millisecond)

	// 10% 的失败概率
	if rand.Float32() < 0.1 {
		return "", errors.New("支付平台处理失败, 请稍后重试(10%概率失败)")
	}

	// 如果请求中已有 transaction_id，则复用，否则生成新的
	transactionID := req.TransactionId
	if transactionID == "" {
		transactionID = generateTransactionID()
	}
	return transactionID, nil
}

// generateTransactionID 生成模拟交易号
func generateTransactionID() string {
	now := time.Now().UnixNano()
	rand.Seed(now)
	randomPart := rand.Intn(1000000)
	return fmt.Sprintf("TXN%d%06d", now, randomPart)
}

// recordPayment 记录支付流水到数据库
func (s *ChargeService) recordPayment(req *pbpayment.ChargeReq, transactionID, status, errorMsg string) error {
	record := model.PaymentRecord{
		TransactionID: transactionID,
		Amount:        req.FinalPrice,
		Status:        status, // 可能是 "SUCCESS" 或 "FAILED"
		ErrorMessage:  errorMsg,
		CreatedAt:     time.Now(),
	}

	// 写入数据库
	if err := dal.DB.Create(&record).Error; err != nil {
		return fmt.Errorf("记录支付流水失败：%v", err)
	}
	return nil
}
