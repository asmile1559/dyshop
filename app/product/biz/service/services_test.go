package service

import (
	"context"
	"testing"

	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	rpcclient "github.com/asmile1559/dyshop/app/frontend/rpc"
	"github.com/asmile1559/dyshop/app/product/biz/dal"
	"github.com/asmile1559/dyshop/app/product/biz/model"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	pb "github.com/asmile1559/dyshop/pb/frontend/product_page"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func init() {
	rpcclient.InitRPCClient()
	err := loadConfig()
	if err != nil {
		return
	}
	dsn := viper.GetString("mysql.dsn")
	print(dsn)
	if err := dal.InitDB(); err != nil {
		logrus.Fatalf("failed to init db: %v", err)
	}
}
func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	return viper.ReadInConfig()
}

// 测试初始化
func setupTestDB(t *testing.T) *gorm.DB {
	dsn := viper.GetString("mysql.dsn")
	print(dsn)
	if err := dal.InitDB(); err != nil {
		logrus.Fatalf("failed to init db: %v", err)
	}
	return dal.DB.Begin()
}

// TestCreateProductService 测试创建产品服务
func TestCreateProductService_Run(t *testing.T) {
	ctx := context.Background()

	t.Run("成功创建产品", func(t *testing.T) {
		// 1. 初始化 RPC 客户端（已通过 rpcclient.InitRPCClient() 完成）
		setupTestDB(t)
		svc := service.NewCreateProductService(ctx)
		req := &pb.CreateProductReq{
			Name:        "Test Product",
			Description: "Test Description",
			Price:       99.99,
			Amount:      7,
		}

		// 2. 调用前端服务（实际触发 RPC 请求到后端）
		resp, err := svc.Run(req)
		require.NoError(t, err)
		print(resp)
		//require.True(t, resp.Success)
	})

	// 其他测试用例（参数校验、错误模拟等）类似调整

	t.Run("参数验证失败", func(t *testing.T) {
		tests := []struct {
			name string
			req  *pbproduct.CreateProductReq
			code codes.Code
		}{
			{
				name: "空产品名称",
				req:  &pbproduct.CreateProductReq{Price: 100},
				code: codes.InvalidArgument,
			},
			{
				name: "无效价格",
				req:  &pbproduct.CreateProductReq{Name: "Test", Price: -1},
				code: codes.InvalidArgument,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				svc := NewCreateProductService(ctx)
				_, err := svc.Run(tt.req)
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.code, st.Code())
			})
		}
	})

	t.Run("数据库错误", func(t *testing.T) {
		// 模拟数据库错误
		//origCreate := model.CreateOrUpdateProduct
		//defer func() { model.CreateOrUpdateProduct = origCreate }()
		//model.CreateOrUpdateProduct = func(db *gorm.DB, p *model.Product) error {
		//	return errors.New("database error")
		//}

		svc := NewCreateProductService(ctx)
		req := &pbproduct.CreateProductReq{
			Name:  "Test",
			Price: 100,
		}

		_, err := svc.Run(req)
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.Internal, st.Code())
	})
}

// TestModifyProductService 测试修改产品服务
func TestModifyProductService_Run(t *testing.T) {
	ctx := context.Background()
	t.Run("成功更新产品", func(t *testing.T) {
		setupTestDB(t)
		// 先创建测试产品
		testProduct := &model.Product{
			Name:  "Original Name",
			Price: 50,
		}
		require.NoError(t, model.CreateOrUpdateProduct(dal.DB, testProduct))

		svc := NewModifyProductService(ctx)
		newName := "Updated Name"
		req := &pbproduct.ModifyProductReq{
			Id:   uint32(testProduct.ID),
			Name: &newName,
		}

		resp, err := svc.Run(req)
		require.NoError(t, err)
		require.True(t, resp.Success)
	})

	t.Run("参数验证失败", func(t *testing.T) {
		tests := []struct {
			name string
			req  *pbproduct.ModifyProductReq
			code codes.Code
		}{
			{
				name: "缺少产品ID",
				req:  &pbproduct.ModifyProductReq{},
				code: codes.InvalidArgument,
			},
			{
				name: "无更新字段",
				req:  &pbproduct.ModifyProductReq{Id: 1},
				code: codes.InvalidArgument,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				svc := NewModifyProductService(ctx)
				_, err := svc.Run(tt.req)
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.code, st.Code())
			})
		}
	})

	t.Run("产品不存在", func(t *testing.T) {
		setupTestDB(t)
		svc := service.NewModifyProductService(ctx)
		name := "New Name"
		req := &pb.ModifyProductReq{
			Id:   999,
			Name: &name,
		}

		_, err := svc.Run(req)
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.NotFound, st.Code())
	})

	t.Run("数据库错误", func(t *testing.T) {
		// 模拟数据库错误
		//origUpdate := model.CreateOrUpdateProduct
		//defer func() { model.CreateOrUpdateProduct = origUpdate }()
		//model.CreateOrUpdateProduct = func(db *gorm.DB, p *model.Product) error {
		//	return errors.New("database error")
		//}

		svc := NewModifyProductService(ctx)
		name := "New Name"
		req := &pbproduct.ModifyProductReq{
			Id:   1,
			Name: &name,
		}

		_, err := svc.Run(req)
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.Internal, st.Code())
	})
}

// TestDeleteProductService 测试删除产品服务
func TestDeleteProductService_Run(t *testing.T) {
	ctx := context.Background()
	db := setupTestDB(t)
	print(db)
	t.Run("成功删除产品", func(t *testing.T) {
		setupTestDB(t)
		// 先创建测试产品
		testProduct := &model.Product{Name: "Test Product"}
		require.NoError(t, model.CreateOrUpdateProduct(dal.DB, testProduct))

		svc := service.NewDeleteProductService(ctx)
		req := &pb.DeleteProductReq{Id: uint32(testProduct.ID)}

		resp, err := svc.Run(req)
		print(resp)
		require.NoError(t, err)
		//require.True(t, resp.Success)
	})

	t.Run("参数验证失败", func(t *testing.T) {
		svc := NewDeleteProductService(ctx)
		_, err := svc.Run(&pbproduct.DeleteProductReq{})
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, st.Code())
	})

	t.Run("产品不存在", func(t *testing.T) {
		setupTestDB(t)
		svc := NewDeleteProductService(ctx)
		req := &pbproduct.DeleteProductReq{Id: 999}

		_, err := svc.Run(req)
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.NotFound, st.Code())
	})

	t.Run("数据库错误", func(t *testing.T) {
		// 模拟数据库错误
		//origDelete := model.DeleteProduct
		//defer func() { model.DeleteProduct = origDelete }()
		//model.DeleteProduct = func(db *gorm.DB, id uint32) error {
		//	return errors.New("database error")
		//}

		svc := NewDeleteProductService(ctx)
		req := &pbproduct.DeleteProductReq{Id: 1}

		_, err := svc.Run(req)
		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.Internal, st.Code())
	})
}
