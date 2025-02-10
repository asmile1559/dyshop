package service

import (
	"context"
	"testing"

	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

func TestGetCartService(t *testing.T) {
	ctx := context.Background()
	userID := uint32(1003)

	emptySrv := NewEmptyCartService(ctx)
	addSrv := NewAddItemService(ctx)
	getSrv := NewGetCartService(ctx)

	// 先清空，确保没有残留
	_, _ = emptySrv.Run(&pbcart.EmptyCartReq{UserId: userID})

	// 第一次获取，应该为空
	resp1, err := getSrv.Run(&pbcart.GetCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("GetCart error: %v", err)
	}
	if len(resp1.Cart.Items) != 0 {
		t.Fatalf("expected empty cart, got %d items", len(resp1.Cart.Items))
	}

	// 添加一条商品
	_, err = addSrv.Run(&pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 999,
			Quantity:  10,
		},
	})
	if err != nil {
		t.Fatalf("AddItem error: %v", err)
	}

	// 再次获取，应该有1条商品
	resp2, err := getSrv.Run(&pbcart.GetCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("GetCart error: %v", err)
	}
	if len(resp2.Cart.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp2.Cart.Items))
	}
	if resp2.Cart.Items[0].ProductId != 999 || resp2.Cart.Items[0].Quantity != 10 {
		t.Fatalf("expected (productId=999, quantity=10), got (productId=%d, quantity=%d)",
			resp2.Cart.Items[0].ProductId, resp2.Cart.Items[0].Quantity)
	}
}
