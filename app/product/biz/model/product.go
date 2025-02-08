package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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

// TableName 自定义表名（proto 服务名为 ProductCatalogService）
func (Product) TableName() string {
	return "product_catalog"
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
