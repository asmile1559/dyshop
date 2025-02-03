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
		return nil, err
	}

	// 2. 模拟调用第三方支付平台进行扣款
	transactionID, err := s.callPaymentGateway(req)
	if err != nil {
		// 若支付失败，可考虑记录失败流水（此处省略）
		return nil, err
	}

	// 3. 支付成功后，记录支付流水到数据库
	if err := s.recordPayment(req, transactionID); err != nil {
		// 若记录流水失败，根据业务逻辑可选择回滚或重试，此处直接返回错误
		return nil, err
	}

	return &pbpayment.ChargeResp{TransactionId: transactionID}, nil
}

// validateRequest 对请求参数进行基本的校验
func (s *ChargeService) validateRequest(req *pbpayment.ChargeReq) error {
	if req == nil {
		return errors.New("请求不能为空")
	}
	if req.Amount <= 0 {
		return errors.New("支付金额必须大于0")
	}
	if req.OrderId == "" {
		return errors.New("订单号不能为空")
	}
	if req.UserId == 0 {
		return errors.New("用户ID不能为空")
	}

	// 校验信用卡信息
	cc := req.CreditCard
	if cc == nil {
		return errors.New("信用卡信息不能为空")
	}
	// 简单校验：卡号非空且长度在合理范围内（例如13~19位）
	ccNumber := strings.TrimSpace(cc.CreditCardNumber)
	if len(ccNumber) < 13 || len(ccNumber) > 19 {
		return errors.New("无效的信用卡号码")
	}
	// 校验 CVV，一般为3或4位数字
	cvvStr := strconv.Itoa(int(cc.CreditCardCvv))
	if len(cvvStr) < 3 || len(cvvStr) > 4 {
		return errors.New("无效的信用卡CVV")
	}

	// 校验过期时间：不能早于当前年月
	now := time.Now()
	expYear := int(cc.CreditCardExpirationYear)
	expMonth := int(cc.CreditCardExpirationMonth)
	if expYear < now.Year() || (expYear == now.Year() && expMonth < int(now.Month())) {
		return errors.New("信用卡已过期")
	}

	return nil
}

// callPaymentGateway 模拟调用第三方支付平台进行扣款处理
func (s *ChargeService) callPaymentGateway(req *pbpayment.ChargeReq) (string, error) {
	// 模拟支付耗时
	time.Sleep(500 * time.Millisecond)

	// 模拟 10% 的支付失败概率
	if rand.Float32() < 0.1 {
		return "", errors.New("支付平台处理失败，请稍后重试")
	}

	// 生成一个模拟的交易ID（时间戳 + 随机数）
	transactionID := generateTransactionID()
	return transactionID, nil
}

// generateTransactionID 生成模拟交易号
func generateTransactionID() string {
	now := time.Now().UnixNano()
	rand.Seed(now)
	randomPart := rand.Intn(1000000)
	return fmt.Sprintf("TXN%d%06d", now, randomPart)
}

// recordPayment 记录支付流水到数据库（使用 dal.DB）
func (s *ChargeService) recordPayment(req *pbpayment.ChargeReq, transactionID string) error {
	// 构造支付流水记录（这里使用 dal/model 中定义的 PaymentRecord 模型）
	record := model.PaymentRecord{
		OrderID:       req.OrderId,
		UserID:        req.UserId,
		Amount:        req.Amount,
		TransactionID: transactionID,
		Status:        "SUCCESS",
		CreatedAt:     time.Now(),
	}

	// 将记录插入数据库
	if err := dal.DB.Create(&record).Error; err != nil {
		return fmt.Errorf("记录支付流水失败：%v", err)
	}
	return nil
}
