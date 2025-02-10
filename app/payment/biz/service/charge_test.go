package service_test

import (
    "context"
    
    "testing"

    "github.com/asmile1559/dyshop/app/payment/biz/dal"
    pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
    "github.com/asmile1559/dyshop/app/payment/biz/service"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/DATA-DOG/go-sqlmock"
)

// 测试 validateRequest 方法
func TestValidateRequest(t *testing.T) {
    // 使用 sqlmock 模拟数据库连接
    mockDB, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer mockDB.Close()

    // 初始化 gorm.DB
    db, err := gorm.Open(mysql.New(mysql.Config{
        Conn:                      mockDB,
        SkipInitializeWithVersion: true,
    }), &gorm.Config{})
    assert.NoError(t, err)

    // 替换全局 DB（注意：实际项目中建议使用依赖注入）
    oldDB := dal.DB
    dal.DB = db
    defer func() { dal.DB = oldDB }()

    tests := []struct {
        name       string
        req        *pbpayment.ChargeReq
        wantErr    string
        expectExec bool // 是否期望执行插入操作
    }{
        {
            name:    "空请求",
            req:     nil,
            wantErr: "请求不能为空",
            expectExec: false,
        },
        {
            name: "金额为0",
            req: &pbpayment.ChargeReq{
                Amount:   0,
                OrderId:  "order123",
                UserId:   1,
                CreditCard: &pbpayment.CreditCardInfo{
                    CreditCardNumber:          "4111111111111111",
                    CreditCardCvv:             123,
                    CreditCardExpirationYear:  2025,
                    CreditCardExpirationMonth: 12,
                },
            },
            wantErr: "支付金额必须大于0",
            expectExec: false,
        },
        {
            name: "订单号为空",
            req: &pbpayment.ChargeReq{
                Amount:   100,
                OrderId:  "",
                UserId:   1,
                CreditCard: &pbpayment.CreditCardInfo{
                    CreditCardNumber:          "4111111111111111",
                    CreditCardCvv:             123,
                    CreditCardExpirationYear:  2025,
                    CreditCardExpirationMonth: 12,
                },
            },
            wantErr: "订单号不能为空",
            expectExec: false,
        },
        {
            name: "用户ID为0",
            req: &pbpayment.ChargeReq{
                Amount:   100,
                OrderId:  "order123",
                UserId:   0,
                CreditCard: &pbpayment.CreditCardInfo{
                    CreditCardNumber:          "4111111111111111",
                    CreditCardCvv:             123,
                    CreditCardExpirationYear:  2025,
                    CreditCardExpirationMonth: 12,
                },
            },
            wantErr: "用户ID不能为空",
            expectExec: false,
        },
        {
            name: "信用卡号长度不足",
            req: &pbpayment.ChargeReq{
                Amount:  100,
                OrderId: "order123",
                UserId:  1,
                CreditCard: &pbpayment.CreditCardInfo{
                    CreditCardNumber:          "411111", // 6位
                    CreditCardCvv:             123,
                    CreditCardExpirationYear:  2025,
                    CreditCardExpirationMonth: 12,
                },
            },
            wantErr: "无效的信用卡号码",
            expectExec: false,
        },
        {
            name: "信用卡CVV长度错误",
            req: &pbpayment.ChargeReq{
                Amount:  100,
                OrderId: "order123",
                UserId:  1,
                CreditCard: &pbpayment.CreditCardInfo{
                    CreditCardNumber:          "4111111111111111",
                    CreditCardCvv:             12, // CVV长度不对
                    CreditCardExpirationYear:  2025,
                    CreditCardExpirationMonth: 12,
                },
            },
            wantErr: "无效的信用卡CVV",
            expectExec: false,
        },
        {
            name: "信用卡已过期",
            req: &pbpayment.ChargeReq{
                Amount:  100,
                OrderId: "order123",
                UserId:  1,
                CreditCard: &pbpayment.CreditCardInfo{
                    CreditCardNumber:          "4111111111111111",
                    CreditCardCvv:             123,
                    CreditCardExpirationYear:  2022, // 已经过期
                    CreditCardExpirationMonth: 12,
                },
            },
            wantErr: "信用卡已过期",
            expectExec: false,
        },
        {
            name: "有效请求",
            req: validRequest(),
            wantErr: "",
            expectExec: true,
        },
    }

    svc := service.NewChargeService(context.Background())
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.expectExec {
                // 设置期望的 SQL 操作
                mock.ExpectBegin()
                mock.ExpectExec(`INSERT INTO .*`).WillReturnResult(sqlmock.NewResult(1, 1))
                mock.ExpectCommit()
            } else {
                mock.ExpectBegin()
                mock.ExpectRollback()
            }

            resp, err := svc.Run(tt.req)
            if tt.wantErr != "" {
                assert.Contains(t, err.Error(), tt.wantErr)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, resp)
            }
            assert.NoError(t, mock.ExpectationsWereMet())
        })
    }
}

// 辅助函数：生成有效请求
func validRequest() *pbpayment.ChargeReq {
    return &pbpayment.ChargeReq{
        Amount:  100,
        OrderId: "order123",
        UserId:  1,
        CreditCard: &pbpayment.CreditCardInfo{
            CreditCardNumber:          "4111111111111111",
            CreditCardCvv:             123,
            CreditCardExpirationYear:  2025,
            CreditCardExpirationMonth: 12,
        },
    }
}