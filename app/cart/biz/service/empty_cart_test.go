package service

import (
	"context"
	"testing"

	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
)

func TestEmptyCartService(t *testing.T) {
	ctx := context.Background()
	userID := uint32(1002)

	emptySrv := NewEmptyCartService(ctx)
	addSrv := NewAddItemService(ctx)
	getSrv := NewGetCartService(ctx)

	// 先清空保证干净
	_, _ = emptySrv.Run(&pbcart.EmptyCartReq{UserId: userID})

	// 添加两条商品
	_, err := addSrv.Run(&pbcart.AddItemReq{
		UserId: userID,
		Item:   &pbcart.CartItem{ProductId: 100, Quantity: 1},
	})
	if err != nil {
		t.Fatalf("AddItem error: %v", err)
	}
	_, err = addSrv.Run(&pbcart.AddItemReq{
		UserId: userID,
		Item:   &pbcart.CartItem{ProductId: 200, Quantity: 2},
	})
	if err != nil {
		t.Fatalf("AddItem error: %v", err)
	}

	// 验证现在购物车有2条
	beforeResp, err := getSrv.Run(&pbcart.GetCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("GetCart error before empty: %v", err)
	}
	if len(beforeResp.Cart.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(beforeResp.Cart.Items))
	}

	// 调用 EmptyCart
	_, err = emptySrv.Run(&pbcart.EmptyCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("EmptyCart error: %v", err)
	}

	// 再次检查，应该变空
	afterResp, err := getSrv.Run(&pbcart.GetCartReq{UserId: userID})
	if err != nil {
		t.Fatalf("GetCart error after empty: %v", err)
	}
	if len(afterResp.Cart.Items) != 0 {
		t.Fatalf("expected 0 items, got %d", len(afterResp.Cart.Items))
	}
}
