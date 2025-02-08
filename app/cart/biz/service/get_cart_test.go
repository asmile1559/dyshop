package service

import (
	"context"
	"testing"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

// TestGetCartService 测试获取购物车功能
func TestGetCartService(t *testing.T) {
	userID := uint32(6666)

	// 开始前先清空，避免受其他测试影响
	dal.ClearCart(userID)

	// 1) 先获取一个空购物车
	getSrv := NewGetCartService(context.Background())
	resp, err := getSrv.Run(&pbcart.GetCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("GetCartService.Run() error: %v", err)
	}
	if resp.Cart == nil {
		t.Fatal("GetCartService.Run() returned nil cart")
	}
	if resp.Cart.UserId != userID {
		t.Fatalf("expected cart.UserId = %d, got %d", userID, resp.Cart.UserId)
	}
	if len(resp.Cart.Items) != 0 {
		t.Fatalf("expected empty cart, got %d item(s)", len(resp.Cart.Items))
	}

	// 2) 添加一条商品后，再获取
	addSrv := NewAddItemService(context.Background())
	_, _ = addSrv.Run(&pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 101,
			Quantity:  1,
		},
	})

	resp2, err := getSrv.Run(&pbcart.GetCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("GetCartService.Run() second call error: %v", err)
	}
	if len(resp2.Cart.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp2.Cart.Items))
	}
	if resp2.Cart.Items[0].ProductId != 101 || resp2.Cart.Items[0].Quantity != 1 {
		t.Fatalf(
			"expected {productId=101, quantity=1}, got {productId=%d, quantity=%d}",
			resp2.Cart.Items[0].ProductId, resp2.Cart.Items[0].Quantity,
		)
	}
}
