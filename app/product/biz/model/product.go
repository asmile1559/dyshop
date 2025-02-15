package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asmile1559/dyshop/pb/backend/product"
	"gorm.io/gorm"
)

// Product 数据库模型，与 Proto 定义对齐
type Product struct {
	gorm.Model
	// 与 proto 中 uint32 id=1 对应
	// 使用自定义主键字段名（默认 gorm.Model 的 ID 是 uint，proto 中为 uint32）
	ID          uint       `gorm:"primaryKey;autoIncrement;column:id"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Description string     `gorm:"type:text"`
	Picture     string     `gorm:"type:varchar(512)"`
	Price       float32    `gorm:"type:decimal(10,2);precision:10;scale:2"` // 对应 proto 的 float
	Categories  Categories `gorm:"type:json"`                               // 处理重复字符串的 JSON 存储
}

func (p *Product) Error() string {
	//TODO implement me
	panic("implement me")
}

// Categories 自定义类型处理 JSON 数组
type Categories []string

// Scan 实现数据库读取时的反序列化
func (c *Categories) Scan(value interface{}) error {
	if value == nil {
		*c = []string{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid database type for categories")
	}
	return json.Unmarshal(bytes, c)
}

// Value 实现数据库写入时的序列化
func (c Categories) Value() (driver.Value, error) {
	if len(c) == 0 {
		return "[]", nil // 避免存储 NULL
	}
	return json.Marshal(c)
}

// 根据产品ID获取产品信息
// 如果不存在则返回 nil
func GetProductByID(db *gorm.DB, productID uint) (*Product, error) {
	var product Product
	err := db.Where("id = ?", productID).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &product, nil
}

// 创建或更新产品信息（使用 Upsert 语义）
func CreateOrUpdateProduct(db *gorm.DB, product *Product) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 尝试查找现有产品
		var existingProduct Product
		err := tx.Where("id = ?", product.ID).First(&existingProduct).Error

		// 处理查找结果
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新记录
			if err := tx.Create(product).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		} else {
			// 更新现有记录（保留原始创建时间）
			product.CreatedAt = existingProduct.CreatedAt
			if err := tx.Save(product).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// 清空所有产品记录（谨慎使用）
func ClearAllProducts(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// 执行物理删除（忽略软删除标记）
		if err := tx.Unscoped().Where("1 = 1").Delete(&Product{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// DeleteProduct 删除产品
func DeleteProduct(db *gorm.DB, id uint32) error {
	result := db.Delete(&Product{}, id)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

// ListProducts 分页查询商品
// 返回值：产品列表，总数，错误信息
func ListProducts(db *gorm.DB, page int32, pageSize int32, category string) ([]*Product, int64, error) {
	const (
		defaultPage     = 1
		defaultPageSize = 20
		maxPageSize     = 100
	)

	// 参数校验与修正
	if page < 1 {
		page = defaultPage
	}
	if pageSize < 1 || pageSize > maxPageSize {
		pageSize = defaultPageSize
	}

	var (
		products []*Product
		total    int64
		query    = db
	)

	// 分类过滤（MySQL JSON_CONTAINS 语法）
	if category != "" {
		query = query.Where("JSON_CONTAINS(categories, ?)", fmt.Sprintf(`"%s"`, category))
	}

	// 获取总数（在分页前）
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 计算分页偏移量（防止溢出）
	offset := (int(page) - 1) * int(pageSize)
	if offset < 0 {
		offset = 0
	}

	// 执行分页查询
	err := query.Order("id DESC").
		Offset(offset).
		Limit(int(pageSize)).
		Find(&products).Error

	if err != nil {
		return nil, 0, err
	}

	return products, total, err
}

// ToProto 转换数据库模型到 Protobuf 结构（用于 service 层）
func (p *Product) ToProto() *product.Product {
	return &product.Product{
		Id:          uint32(p.ID), // uint32 类型转换
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Picture,
		Price:       p.Price,
		Categories:  p.Categories,
	}
}

// FromProto 从 Protobuf 结构转换（用于创建/更新操作）
func (p *Product) FromProto(protoProduct *product.Product) {
	p.ID = uint(protoProduct.Id) // 注意类型转换
	p.Name = protoProduct.Name
	p.Description = protoProduct.Description
	p.Picture = protoProduct.Picture
	p.Price = protoProduct.Price
	p.Categories = protoProduct.Categories
}
