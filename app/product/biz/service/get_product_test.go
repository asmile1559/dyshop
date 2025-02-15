package service

import (
	"testing"
)

func TestProductService_GetProduct(t *testing.T) {
	//ctx := context.Background()
	//svc := service.NewCreateProductService(ctx)
	//
	//// 先创建测试产品
	//createReq := &product_page.CreateProductReq{
	//	Name:        "Test Product",
	//	Description: "Test Description",
	//	Price:       99.99,
	//	Amount:      10,
	//}
	//createResp, err := svc.Run(createReq)
	//require.NoError(t, err)
	//
	//t.Run("成功获取产品", func(t *testing.T) {
	//	req := &product_page.GetProductReq{
	//		Id: createResp., // 假设返回中带有产品ID
	//	}
	//
	//	resp, err := svc.GetProduct(req)
	//	require.NoError(t, err)
	//	require.Equal(t, "Test Product", resp.Product.Name)
	//})
	//
	//t.Run("获取不存在的产品", func(t *testing.T) {
	//	req := &product_page.GetProductReq{
	//		Id: 9999,
	//	}
	//
	//	_, err := svc.GetProduct(req)
	//	require.Error(t, err)
	//	st, ok := status.FromError(err)
	//	require.True(t, ok)
	//	require.Equal(t, codes.NotFound, st.Code())
	//})
}
