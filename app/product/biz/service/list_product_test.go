package service

import (
	"context"
	"testing"

	"github.com/asmile1559/dyshop/app/frontend/biz/service"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestListProductService_Run(t *testing.T) {
	// 初始化基础上下文
	ctx := context.Background()
	t.Run("正常分页查询", func(t *testing.T) {

		// 执行测试
		svc := service.NewListProductService(ctx)
		result, err := svc.Run(&pbproduct.ListProductsReq{
			Page:         1,
			PageSize:     10,
			CategoryName: "drink",
		})

		// 验证结果
		require.NoError(t, err)
		require.Contains(t, result, "resp")

		resp, ok := result["resp"].(*pbproduct.ListProductsResp)
		require.True(t, ok)
		//require.Equal(t, int64(100), resp.Total)
		require.Len(t, resp.Products, 1)
	})

	t.Run("参数验证失败", func(t *testing.T) {
		// 不需要设置 Mock，因为参数验证应该在调用 RPC 前失败
		svc := service.NewListProductService(ctx)

		// 测试用例组
		tests := []struct {
			name string
			req  *pbproduct.ListProductsReq
			code codes.Code
		}{
			{
				name: "页码过小",
				req:  &pbproduct.ListProductsReq{Page: 0},
				code: codes.InvalidArgument,
			},
			{
				name: "分页尺寸过大",
				req:  &pbproduct.ListProductsReq{PageSize: 200},
				code: codes.InvalidArgument,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := svc.Run(tt.req)
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.code, st.Code())
			})
		}
	})

	t.Run("RPC调用失败", func(t *testing.T) {

		svc := service.NewListProductService(ctx)
		_, err := svc.Run(&pbproduct.ListProductsReq{
			Page:     1,
			PageSize: 10,
		})

		require.Error(t, err)
		st, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.Internal, st.Code())
	})

	t.Run("空结果处理", func(t *testing.T) {

		svc := service.NewListProductService(ctx)
		result, err := svc.Run(&pbproduct.ListProductsReq{
			Page:     1,
			PageSize: 10,
		})

		require.NoError(t, err)
		resp := result["resp"].(*pbproduct.ListProductsResp)
		//require.Equal(t, int64(0), resp)
		require.Empty(t, resp.Products)
	})
}
