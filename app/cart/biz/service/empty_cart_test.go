package service

import (
	"context"
	"testing"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

// TestEmptyCartService 测试清空购物车功能
func TestEmptyCartService(t *testing.T) {
	userID := uint32(7777)

	// 开始前先清空，避免受其他测试影响
	dal.ClearCart(userID)

	// 先给购物车添加几条商品
	addSrv := NewAddItemService(context.Background())
	_, _ = addSrv.Run(&pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 111,
			Quantity:  2,
		},
	})
	_, _ = addSrv.Run(&pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 222,
			Quantity:  3,
		},
	})

	// 此时确认购物车不为空
	c := dal.GetCartByUserID(userID)
	if len(c.Items) == 0 {
		t.Fatal("expected cart to have items, but it is empty")
	}

	// 调用 EmptyCart
	srv := NewEmptyCartService(context.Background())
	if _, err := srv.Run(&pbcart.EmptyCartReq{UserId: userID}); err != nil {
		t.Fatalf("EmptyCartService.Run() error: %v", err)
	}

	// 再次检查购物车，应该已清空
	c2 := dal.GetCartByUserID(userID)
	if len(c2.Items) != 0 {
		t.Fatalf("expected empty cart, got %v", c2.Items)
	}
}
