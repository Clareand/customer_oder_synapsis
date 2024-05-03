package model

type Products struct {
	ProductID string `json:"product_id" gorm:"column:product_id"`
	Name      string `json:"name" gorm:"column:name"`
	Stock     int    `json:"stock" gorm:"column:stock"`
	Price     int    `json:"price" gorm:"column:price"`
	Category  string `json:"category" gorm:"column:category"`
}
