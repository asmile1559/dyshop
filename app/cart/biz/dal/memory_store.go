package dal

import (
	"sync"

	"github.com/asmile1559/dyshop/app/cart/biz/model"
)

var (
	cartStore = make(map[uint32]*model.Cart)
	mu        sync.RWMutex
)

// GetCartByUserID 获取指定用户的购物车，如果没有则创建空购物车
func GetCartByUserID(userID uint32) *model.Cart {
	mu.RLock()
	cart, ok := cartStore[userID]
	mu.RUnlock()

	if !ok {
		// 不存在则初始化
		cart = &model.Cart{
			UserID: userID,
			Items:  []model.CartItem{},
		}
		// 存储
		mu.Lock()
		cartStore[userID] = cart
		mu.Unlock()
	}

	return cart
}

// SaveCart 保存购物车到内存 map
func SaveCart(cart *model.Cart) {
	mu.Lock()
	defer mu.Unlock()
	cartStore[cart.UserID] = cart
}

// ClearCart 清空购物车
func ClearCart(userID uint32) {
	mu.Lock()
	defer mu.Unlock()
	cartStore[userID] = &model.Cart{
		UserID: userID,
		Items:  []model.CartItem{},
	}
}
