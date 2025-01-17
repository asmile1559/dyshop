package db

var curTables = []any{
	User{},
	Product{},
	Order{},
}

// TODO: table的抽象结构
type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Role     string
}

type Product struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	Price       int
	Description string
}

type Order struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	ProductID int
	Amount    int
}
