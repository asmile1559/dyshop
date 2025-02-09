package dal

import (
	"errors"

	"github.com/asmile1559/dyshop/app/cart/biz/model"
	"gorm.io/gorm"
)

// 查找指定用户的购物车，不存在则新建一条
func GetOrCreateCartByUserID(userID uint64) (*model.Cart, error) {
	var cart model.Cart
	err := DB.Where("user_id = ?", userID).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在则初始化一个
		cart = model.Cart{
			UserId: userID,
		}
		if errCreate := DB.Create(&cart).Error; errCreate != nil {
			return nil, errCreate
		}
		return &cart, nil
	} else if err != nil {
		return nil, err
	}
	return &cart, nil
}

// 读取某用户的 Cart 及其 CartItems
// 如果还没有创建过购物车，则返回 nil
func GetCartByUserID(userID uint64) (*model.Cart, error) {
	var cart model.Cart
	err := DB.Where("user_id = ?", userID).
		Preload("CartItems"). // 预加载关联的 items
		First(&cart).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 用户没有 cart 记录，返回 nil
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &cart, nil
}

// 为指定用户的购物车添加 / 更新商品
func AddOrUpdateCartItem(userID, productID uint64, quantity int) error {
	// 1) 确保用户有 cart
	cart, err := GetOrCreateCartByUserID(userID)
	if err != nil {
		return err
	}

	// 2) 找该 cart 下是否已有此商品
	var item model.CartItem
	err = DB.Where("cart_id = ? AND product_id = ?", cart.ID, productID).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在则插入新纪录
		item = model.CartItem{
			CartId:    cart.ID,
			ProductId: productID,
			Quantity:  quantity,
		}
		return DB.Create(&item).Error
	} else if err != nil {
		return err
	}

	// 如果已存在，则累加数量
	item.Quantity += quantity
	return DB.Save(&item).Error
}

// 清空指定用户的购物车(删除所有条目)
func ClearCart(userID uint64) error {
	// 先找到 cart
	var cart model.Cart
	err := DB.Where("user_id = ?", userID).First(&cart).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有创建过cart，等于已经是空的
		return nil
	} else if err != nil {
		return err
	}
	// 删除 cart 下的所有 item
	return DB.Where("cart_id = ?", cart.ID).Delete(&model.CartItem{}).Error
}
