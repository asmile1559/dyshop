package dal

import (
	"errors"

	"github.com/asmile1559/dyshop/app/cart/biz/model"
	"gorm.io/gorm"
)

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
// 这里使用 DB.Transaction 保证操作的原子性
func AddOrUpdateCartItem(userID, productID uint64, quantity int) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 1) 在事务内找 cart
		var cart model.Cart
		err := tx.Where("user_id = ?", userID).First(&cart).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在则初始化一个
			cart = model.Cart{
				UserId: userID,
			}
			if errCreate := tx.Create(&cart).Error; errCreate != nil {
				return errCreate // 会回滚事务
			}
		} else if err != nil {
			return err // 会回滚事务
		}

		// 2) 找该 cart 下是否已有此商品
		var item model.CartItem
		err = tx.Where("cart_id = ? AND product_id = ?", cart.ID, productID).
			First(&item).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在则插入新纪录
			item = model.CartItem{
				CartId:    cart.ID,
				ProductId: productID,
				Quantity:  quantity,
			}
			if errCreateItem := tx.Create(&item).Error; errCreateItem != nil {
				return errCreateItem // 回滚
			}
		} else if err != nil {
			return err // 回滚
		} else {
			// 如果已存在，则累加数量
			item.Quantity += quantity
			if errSave := tx.Save(&item).Error; errSave != nil {
				return errSave // 回滚
			}
		}

		// 若整个流程都成功执行，事务会在函数return nil时自动commit
		return nil
	})
}

// 清空指定用户的购物车(删除所有条目)
func ClearCart(userID uint64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 1) 找 cart
		var cart model.Cart
		err := tx.Where("user_id = ?", userID).First(&cart).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有创建过cart，等于已经是空的
			return nil
		} else if err != nil {
			return err
		}

		// 2) 删除 cart 下的所有 item
		if errDel := tx.Where("cart_id = ?", cart.ID).
			Delete(&model.CartItem{}).Error; errDel != nil {
			return errDel
		}

		// 如果把 cart 这一行也删掉(彻底释放Cart)
		// if errDelCart := tx.Delete(&cart).Error; errDelCart != nil {
		// 	return errDelCart
		// }

		return nil
	})
}
