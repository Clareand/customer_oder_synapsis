package model

type Carts struct {
	CartID   string `json:"cart_id" gorm:"column:cart_id"`
	Name     string `json:"name" gorm:"column:name"`
	Quantity string `json:"quantity" gorm:"column:quantity"`
}

type CartItems struct {
	CartItemID     string `json:"cart_item_id" gorm:"column:cart_item_id"`
	ProductID      string `json:"product_id" gorm:"column:product_id"`
	ProductName    string `json:"product_name" gorm:"column:product_name"`
	Quantity       string `json:"quantity" gorm:"column:quantity"`
	CheckoutStatus string `json:"checkout_status" gorm:"column:checkout_status"`
}

type AddToCart struct {
	CartID     string `json:"cart_id" gorm:"column:cart_id"`
	CustomerID string `json:"customer_id" gorm:"column:customer_id"`
	ProductID  string `json:"product_id" gorm:"column:product_id"`
	Quantity   int    `json:"quantity" gorm:"column:quantity"`
}
