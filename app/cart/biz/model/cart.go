package model

import (
	"errors"

	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	UID       uint64
	ProductId uint64 // 商品ID
	Quantity  int    // 数量
}

// GetCartItemsByUserID 查询某用户所有 cart_items 条目
// 如果没有找到任何条目，就返回一个空的切片
func GetCartItemsByUserID(db *gorm.DB, uid uint64) ([]CartItem, error) {
	var items []CartItem
	err := db.Where("uid = ?", uid).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

// 为指定用户的购物车添加 / 更新商品
// 这里使用 DB.Transaction 保证操作的原子性
// AddOrUpdateCartItem 在 cart_items 表里，为指定 (UID, ProductId) 添加或更新商品条目：
//   - 如果存在则累加数量
//   - 如果不存在则插入新纪录
func AddOrUpdateCartItem(db *gorm.DB, uid, productID uint64, quantity int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var item CartItem
		err := tx.Where("uid = ? AND product_id = ?", uid, productID).
			First(&item).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 新建
			item = CartItem{
				UID:       uid,
				ProductId: productID,
				Quantity:  quantity,
			}
			if errCreate := tx.Create(&item).Error; errCreate != nil {
				return errCreate
			}
		} else if err != nil {
			return err
		} else {
			// 已存在 => 累加
			item.Quantity += quantity
			if item.Quantity < 0 {
				item.Quantity = 0
			}
			if errSave := tx.Save(&item).Error; errSave != nil {
				return errSave
			}
		}
		return nil
	})
}

// DeleteItems 根据传入的 item IDs 删除指定用户的部分条目
func DeleteItems(db *gorm.DB, uid uint64, itemIDs []uint64) error {
	if len(itemIDs) == 0 {
		return nil
	}
	return db.Where("uid = ? AND id IN ?", uid, itemIDs).Delete(&CartItem{}).Error
}

// ClearAllItems 删除某用户的所有条目
func ClearAllItems(db *gorm.DB, uid uint64) error {
	return db.Where("uid = ?", uid).Delete(&CartItem{}).Error
}
