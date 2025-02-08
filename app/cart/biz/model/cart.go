package model

type CartItem struct {
	ProductID uint32
	Quantity  int32
}

type Cart struct {
	UserID uint32
	Items  []CartItem
}
