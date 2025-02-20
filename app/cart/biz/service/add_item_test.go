package service

import (
	"context"
	"testing"

	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

func TestAddItemService(t *testing.T) {
	ctx := context.Background()
	userID := uint32(1001)

	// 先清空用户的购物车，确保不受其他测试干扰
	_, _ = NewEmptyCartService(ctx).Run(&pbcart.EmptyCartReq{
		UserId: userID,
	})

	addSrv := NewAddItemService(ctx)
	getSrv := NewGetCartService(ctx)

	// 1) 第一次添加商品
	_, err := addSrv.Run(&pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 123,
			Quantity:  5,
		},
	})
	if err != nil {
		t.Fatalf("AddItemService Run error: %v", err)
	}

	resp1, err := getSrv.Run(&pbcart.GetCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("GetCartService Run error after add 123: %v", err)
	}
	if len(resp1.Cart.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp1.Cart.Items))
	}
	if resp1.Cart.Items[0].ProductId != 123 || resp1.Cart.Items[0].Quantity != 5 {
		t.Fatalf("expected product=123, quantity=5, got product=%d, quantity=%d",
			resp1.Cart.Items[0].ProductId, resp1.Cart.Items[0].Quantity)
	}

	// 2) 再添加相同的商品
	_, err = addSrv.Run(&pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 123,
			Quantity:  3,
		},
	})
	if err != nil {
		t.Fatalf("AddItemService second run error: %v", err)
	}

	resp2, err := getSrv.Run(&pbcart.GetCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("GetCartService Run error after second add: %v", err)
	}
	if len(resp2.Cart.Items) != 1 {
		t.Fatalf("expected still 1 item, got %d", len(resp2.Cart.Items))
	}
	if resp2.Cart.Items[0].Quantity != 8 {
		t.Fatalf("expected quantity=8 (5+3), got %d", resp2.Cart.Items[0].Quantity)
	}

	// 3) 再添加另一个商品
	_, err = addSrv.Run(&pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 999,
			Quantity:  2,
		},
	})
	if err != nil {
		t.Fatalf("AddItemService third run error: %v", err)
	}

	resp3, err := getSrv.Run(&pbcart.GetCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("GetCartService Run error after third add: %v", err)
	}
	if len(resp3.Cart.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(resp3.Cart.Items))
	}
}
