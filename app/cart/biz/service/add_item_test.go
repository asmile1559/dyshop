package service

import (
	"context"
	"testing"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

// TestAddItemService 测试添加商品到购物车功能
func TestAddItemService(t *testing.T) {
	userID := uint32(9999)

	// 开始前先清空，避免受其他测试影响
	dal.ClearCart(userID)

	ctx := context.Background()
	srv := NewAddItemService(ctx)

	// 1) 测试给空购物车添加一条商品
	req1 := &pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 123,
			Quantity:  5,
		},
	}
	if _, err := srv.Run(req1); err != nil {
		t.Fatalf("AddItemService.Run() error: %v", err)
	}

	cart := dal.GetCartByUserID(userID)
	if len(cart.Items) != 1 {
		t.Fatalf("expected 1 item in cart, got %d", len(cart.Items))
	}
	if cart.Items[0].ProductID != 123 || cart.Items[0].Quantity != 5 {
		t.Fatalf("expected (productID=123, quantity=5), got (productID=%d, quantity=%d)",
			cart.Items[0].ProductID, cart.Items[0].Quantity)
	}

	// 2) 测试再次添加同一个商品，数量应累加
	req2 := &pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 123,
			Quantity:  3,
		},
	}
	if _, err := srv.Run(req2); err != nil {
		t.Fatalf("AddItemService.Run() error on second add: %v", err)
	}

	cart2 := dal.GetCartByUserID(userID)
	if len(cart2.Items) != 1 {
		t.Fatalf("expected still 1 item in cart, got %d", len(cart2.Items))
	}
	if cart2.Items[0].Quantity != 8 {
		t.Fatalf("expected quantity to be 8 (5+3), got %d", cart2.Items[0].Quantity)
	}
}
